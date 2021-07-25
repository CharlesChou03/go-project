package services

import (
	"github.com/CharlesChou03/url-shortening-service.git/config"
	"github.com/CharlesChou03/url-shortening-service.git/internal/db"
	"github.com/CharlesChou03/url-shortening-service.git/models"
	modelsReq "github.com/CharlesChou03/url-shortening-service.git/models/request"
	modelsRes "github.com/CharlesChou03/url-shortening-service.git/models/response"
	"github.com/CharlesChou03/url-shortening-service.git/utils"
)

func GenerateShorteningUrl(req *modelsReq.GenShorteningUrlReq, res *modelsRes.GenShorteningUrlRes) (int64, *modelsRes.GenShorteningUrlRes, models.UrlProcessingError) {
	if !req.Validate() {
		return 400, res, models.BadRequestError
	}
	t := utils.GetCurrentMillisecondTimestamp()
	shorteningUrl := UrlShorten(req.OriginalUrl, config.Env.ShorteningUrlLength)
	urlPrefix := config.Env.UrlHost
	url := urlPrefix + shorteningUrl
	expiredAt := getExpiredAt(req.ExpiredAt, t)

	insertUrlReq := db.ShorteningUrlData{}
	insertUrlReq.UserId = req.UserId
	insertUrlReq.OriginalUrl = req.OriginalUrl
	insertUrlReq.ShorteningUrl = url
	insertUrlReq.ExpiredAt = expiredAt
	insertUrlReq.CreatedAt = t
	dbRes, err := db.UrlDB.CreateUrlData(&insertUrlReq)
	if err != nil {
		return 500, res, models.InternalServerError
	}
	if dbRes == "duplicate key error" {
		return 409, res, models.UrlProcessingError{40901, "shortening url conflict"}
	}

	res.OriginalUrl = req.OriginalUrl
	res.ShorteningUrl = url
	res.ExpiredAt = expiredAt
	return 200, res, models.NoError
}

func getExpiredAt(t int64, now int64) int64 {
	if t == 0 {
		return now + config.Env.DefaultExpiredPeriod
	}
	return t
}

func GetOriginalUrl(req *modelsReq.GetOriginalUrlReq, res *modelsRes.GetOriginalUrlRes) (int64, *modelsRes.GetOriginalUrlRes, models.UrlProcessingError) {
	if !req.Validate() {
		return 400, res, models.BadRequestError
	}

	queryUrlReq := db.QueryUrlData{}
	queryUrlReq.ShorteningUrl = req.ShorteningUrl
	queryUrlReq.ExpiredAtEffectiveStart = utils.GetCurrentMillisecondTimestamp()
	total, resp, err := db.UrlDB.QueryUrlData(&queryUrlReq)
	if err != nil {
		return 500, res, models.InternalServerError
	}
	if total > 1 {
		return 409, res, models.UrlProcessingError{40901, "duplicate original url for one shortening url"}
	}
	if total == 0 {
		return 204, res, models.NotFoundError
	}
	urlData := resp[0]
	res.OriginalUrl = urlData.OriginalUrl
	res.ShorteningUrl = urlData.ShorteningUrl
	res.ExpiredAt = urlData.ExpiredAt
	res.CreatedAt = urlData.CreatedAt

	return 200, res, models.NoError
}
