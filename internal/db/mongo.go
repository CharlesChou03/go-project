package db

import (
	"context"

	"github.com/CharlesChou03/url-shortening-service.git/config"
	"github.com/CharlesChou03/url-shortening-service.git/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

type MongoSavingUrlDB struct {
	Client *mongo.Client
}

var (
	collectionShorteningUrl string = "shorteningUrl"
	databaseSavingUrl       string
	UrlDB                   *MongoSavingUrlDB
)

// SetupMongoDB setup mongo db
func SetupMongoDB() *MongoSavingUrlDB {
	cs, err := connstring.ParseAndValidate(config.Env.MongoURI)
	if err != nil {
		logger.Error.Fatalf("SetupMongoDB parse MongoURI error %+v\n", err)
	}
	databaseSavingUrl = cs.Database
	// Set client options
	clientOptions := options.Client().ApplyURI(config.Env.MongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logger.Error.Fatalf("SetupMongoDB connect error %+v\n", err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		logger.Error.Fatalf("SetupMongoDB ping error %+v\n", err)
	}
	logger.Info.Printf("SetupMongoDB successful")
	return &MongoSavingUrlDB{Client: client}
}

// Close disconnect DB connection
func (c *MongoSavingUrlDB) Close() {
	if err := c.Client.Disconnect(context.TODO()); err != nil {
		logger.Error.Fatalf("MongoSavingUrlDB close error %+v\n", err)
	}
}

func (c *MongoSavingUrlDB) CreateUrlData(r *ShorteningUrlData) (string, error) {
	collection := c.Client.Database(databaseSavingUrl).Collection(collectionShorteningUrl)
	_, err := collection.InsertOne(context.TODO(), r)
	if err != nil {
		errMessage := err.(mongo.WriteException)
		errCode := errMessage.WriteErrors[0].Code
		logger.Error.Printf("Create shorten url error %+v\n", err)
		if errCode == 11000 {
			return "duplicate or error", err
		}
		return "", err
	}
	return "ok", err
}

func (c *MongoSavingUrlDB) QueryUrlData(req *QueryUrlData) (int64, []ShorteningUrlData, error) {
	collection := c.Client.Database(databaseSavingUrl).Collection(collectionShorteningUrl)

	findOptions := options.Find()
	findOptions.SetSkip(req.From)
	findOptions.SetLimit(req.Size)

	var urlDataList = []ShorteningUrlData{}
	var setElements bson.D
	if req.UserId != "" {
		setElements = append(setElements, bson.E{"userId", req.UserId})
	}
	if req.OriginalUrl != "" {
		setElements = append(setElements, bson.E{"originalUrl", req.OriginalUrl})
	}
	if req.ShorteningUrl != "" {
		setElements = append(setElements, bson.E{"shorteningUrl", req.ShorteningUrl})
	}
	if req.ExpiredAtEffectiveStart != 0 {
		setElements = append(setElements, bson.E{"expiredAt", bson.M{"$gte": req.ExpiredAtEffectiveStart}})
	}
	if req.ExpiredAtEffectiveEnd != 0 {
		setElements = append(setElements, bson.E{"expiredAt", bson.M{"$lte": req.ExpiredAtEffectiveEnd}})
	}
	if req.CreatedAtEffectiveStart != 0 {
		setElements = append(setElements, bson.E{"createdAt", bson.M{"$gte": req.CreatedAtEffectiveStart}})
	}
	if req.CreatedAtEffectiveEnd != 0 {
		setElements = append(setElements, bson.E{"createdAt", bson.M{"$lte": req.CreatedAtEffectiveEnd}})
	}

	total, err := collection.CountDocuments(context.TODO(), setElements)
	if err != nil {
		logger.Error.Printf("Count shortening url error %+v\n", err)
		return total, urlDataList, err
	}
	cursor, err := collection.Find(context.TODO(), setElements)
	if err != nil {
		logger.Error.Printf("Get shortening url error %+v\n", err)
		return total, urlDataList, err
	}
	for cursor.Next(context.TODO()) {
		urlData := ShorteningUrlData{}
		err := cursor.Decode(&urlData)
		if err != nil {
			logger.Error.Printf("Get shortening url parse error %+v\n", err)
			return total, urlDataList, err
		}
		urlDataList = append(urlDataList, urlData)
	}
	return total, urlDataList, err
}
