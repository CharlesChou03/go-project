package db

type ShorteningUrlData struct {
	UserId        string `bson:"userId" json:"userId"`
	OriginalUrl   string `bson:"originalUrl" json:"originalUrl"`
	ShorteningUrl string `bson:"shorteningUrl" json:"shorteningUrl"`
	ExpiredAt     int64  `bson:"expiredAt" json:"expiredAt"`
	CreatedAt     int64  `bson:"createdAt" json:"createdAt"`
}

type QueryUrlData struct {
	UserId                  string `bson:"userId" json:"userId"`
	OriginalUrl             string `bson:"originalUrl" json:"originalUrl"`
	ShorteningUrl           string `bson:"shorteningUrl" json:"shorteningUrl"`
	ExpiredAtEffectiveStart int64  `bson:"expiredAtEffectiveStart" json:"expiredAtEffectiveStart"`
	ExpiredAtEffectiveEnd   int64  `bson:"expiredAtEffectiveEnd" json:"expiredAtEffectiveEnd"`
	CreatedAtEffectiveStart int64  `bson:"createdAtEffectiveStart" json:"createdAtEffectiveStart"`
	CreatedAtEffectiveEnd   int64  `bson:"createdAtEffectiveEnd" json:"createdAtEffectiveEnd"`
	From                    int64  `bson:"from" json:"from"`
	Size                    int64  `bson:"size" json:"size"`
}
