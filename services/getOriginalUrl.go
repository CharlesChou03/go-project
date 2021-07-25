package services

import (
	"github.com/CharlesChou03/url-shortening-service.git/internal/db"
	"github.com/CharlesChou03/url-shortening-service.git/models"
	modelsReq "github.com/CharlesChou03/url-shortening-service.git/models/request"
	modelsRes "github.com/CharlesChou03/url-shortening-service.git/models/response"
	"github.com/CharlesChou03/url-shortening-service.git/utils"
)

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
