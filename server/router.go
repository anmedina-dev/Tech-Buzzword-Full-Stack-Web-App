package server

import (
	"os"
	"strings"
	"tech-buzzword-service/controllers"
	"tech-buzzword-service/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	// Comment this when running locally
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	whitelistedIPs := strings.Split(os.Getenv("WHITELISTED_IPS"), ",")
	router.ForwardedByClientIP = true
	router.SetTrustedProxies(whitelistedIPs)

	health := new(controllers.HealthController)

	router.GET(os.Getenv("HEALTH"), health.Status)

	router.Use(middleware.AuthMiddleware())

	version := router.Group(os.Getenv("VERSION"))
	{
		buzzword := new(controllers.BuzzwordController)
		version.GET(os.Getenv("BUZZWORD_ROUTE"), buzzword.RetrieveBuzzword)
		version.GET(os.Getenv("PREVIOUS_BUZZWORDS_ROUTE"), buzzword.RetrievePreviousBuzzwords)
		version.POST(os.Getenv("SET_NEW_BUZZWORD"), buzzword.SetNewBuzzword)
	}

	return router

}
