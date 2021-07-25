package handlers

import (
	"net/http"

	"github.com/CharlesChou03/url-shortening-service.git/logger"
	"github.com/CharlesChou03/url-shortening-service.git/models"
	modelsReq "github.com/CharlesChou03/url-shortening-service.git/models/request"
	modelsRes "github.com/CharlesChou03/url-shortening-service.git/models/response"
	"github.com/CharlesChou03/url-shortening-service.git/services"
	"github.com/gin-gonic/gin"
)

// GenerateShorteningUrlHandler create shortening url
// @Summary create shortening url
// @Description create shortening url
// @Tags Shorten Url
// @Accept json
// @Produce json
// @Param body body modelsReq.GenShorteningUrlReq true "body"
// @Success 200 {object} models.GenShorteningUrlRes "ok"
// @Failure 400 {object} models.UrlProcessingError "bad request"
// @Failure 409 {object} models.UrlProcessingError "shortening url conflict"
// @Failure 500 {object} models.UrlProcessingError "internal server error"
// @Router /api/url-shortening-service/v1/generate [post]
func GenerateShorteningUrlHandler(c *gin.Context) {
	req := modelsReq.GenShorteningUrlReq{}
	res := modelsRes.GenShorteningUrlRes{}
	c.BindJSON(&req)
	logger.Info.Printf("[GenerateShorteningUrlHandler] req=%+v\n", req)
	statusCode, resp, err := services.GenerateShorteningUrl(&req, &res)
	switch statusCode {
	case 200:
		c.JSON(http.StatusOK, resp)
	case 400:
		c.JSON(http.StatusBadRequest, err)
	case 409:
		c.JSON(http.StatusFailedDependency, err)
	case 500:
		c.JSON(http.StatusInternalServerError, err)
	default:
		c.JSON(http.StatusInternalServerError, models.InternalServerError)
	}
}

// GetOriginalUrlHandler get original url by shortening url
// @Summary get original url by shortening url
// @Description get original url by shortening url
// @Tags Shorten Url
// @Accept json
// @Produce json
// @Param body body modelsReq.GetOriginalUrlReq true "body"
// @Success 200 {object} models.GetOriginalUrlRes "ok"
// @Failure 204 {object} models.UrlProcessingError "url not found"
// @Failure 400 {object} models.UrlProcessingError "bad request"
// @Failure 409 {object} models.UrlProcessingError "shortening url conflict"
// @Failure 500 {object} models.UrlProcessingError "internal server error"
// @Router /api/url-shortening-service/v1/getoriginalurl [post]
func GetOriginalUrlHandler(c *gin.Context) {
	req := modelsReq.GetOriginalUrlReq{}
	res := modelsRes.GetOriginalUrlRes{}
	c.BindJSON(&req)
	logger.Info.Printf("[GetOriginalUrlHandler] req=%+v\n", req)
	statusCode, resp, err := services.GetOriginalUrl(&req, &res)
	switch statusCode {
	case 200:
		c.JSON(http.StatusOK, resp)
	case 204:
		c.JSON(http.StatusNoContent, err)
	case 400:
		c.JSON(http.StatusBadRequest, err)
	case 409:
		c.JSON(http.StatusConflict, err)
	case 500:
		c.JSON(http.StatusInternalServerError, err)
	default:
		c.JSON(http.StatusInternalServerError, models.InternalServerError)
	}
}
