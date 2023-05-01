package routes

import (
	"tranphrase/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.AuthMiddleware())

	api := router.Group("/api")
	{
		TranslateRoutes(api)
	}

	return router
}
