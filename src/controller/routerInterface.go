package controller

import "github.com/gin-gonic/gin"

type IRouter interface {
	GetRoutes(appGroup *gin.RouterGroup)
	InitSubRoutes(appGroup *gin.RouterGroup)
}
