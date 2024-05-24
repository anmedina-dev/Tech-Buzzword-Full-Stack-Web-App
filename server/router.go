package server

import (
	"os"
	"tech-buzzword-service/controllers"
	"tech-buzzword-service/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)
	buzzword := new(controllers.Buzzword)

	router.GET(os.Getenv("HEALTH"), health.Status)

	router.Use(middleware.AuthMiddleware())

	version := router.Group(os.Getenv("VERSION"))
	{
		version.GET(os.Getenv("BUZZWORD_ROUTE"), buzzword.Retrieve)
	}

	return router

}
