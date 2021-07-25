package main

import (
	"log"

	"github.com/CharlesChou03/url-shortening-service.git/config"
	_ "github.com/CharlesChou03/url-shortening-service.git/docs"
	"github.com/CharlesChou03/url-shortening-service.git/internal/db"
	"github.com/CharlesChou03/url-shortening-service.git/internal/handlers"
	"github.com/CharlesChou03/url-shortening-service.git/logger"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func setup() {
	logger.Setup()
	config.Setup()
	db.UrlDB = db.SetupMongoDB()
}

func errorHandlingMiddleWare(log *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err == nil {
			return
		}

		log.Printf("unexpected error: %s\n", err.Error())
	}
}

func setupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger(), errorHandlingMiddleWare(logger.Error))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/health", handlers.HealthHandler)
	r.GET("/version", handlers.VersionHandler)
	r.POST("/api/url-shortening-service/v1/generate", handlers.GenerateShorteningUrlHandler)
	r.POST("/api/url-shortening-service/v1/getoriginalurl", handlers.GetOriginalUrlHandler)

	return r
}

// @title Shortening Url Swagger
// @version 0.0.1
// @description this service is for shortening url
func main() {
	setup()
	defer db.UrlDB.Close()
	r := setupRouter()
	r.Run(":9999")
}
