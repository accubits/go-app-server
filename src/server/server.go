package server

import (
	cfg "config"
	"controller"
	"logger"

	"github.com/gin-gonic/gin"
)

var log = logger.NewLogger("APP")

var router *gin.Engine

func Init() {
	config := cfg.GetConfig()
	router = gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	mainCntrlr := new(controller.MainController)
	appVer := config.GetString("app.version")
	if appVer == "" {
		appVer = "v1"
	}
	mainCntrlr.GetRoutes(router.Group(appVer))
	log.Info("server starting at :: ", config.GetString("server.port"))
	router.Run(":" + config.GetString("server.port"))
}

func GetRouter() *gin.Engine {
	return router
}
