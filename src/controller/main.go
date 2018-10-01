package controller

import (
	"controller/user"

	"github.com/gin-gonic/gin"
)

type MainController struct{}

func (mc *MainController) GetRoutes(appGroup *gin.RouterGroup) {
	mc.InitSubRoutes(appGroup)
}

func (mc *MainController) InitSubRoutes(appGroup *gin.RouterGroup) {
	controllerMap := makeControllerMap()
	for key, cntrlr := range controllerMap {
		cntrlr.GetRoutes(appGroup.Group(key))
	}
}

func makeControllerMap() map[string]IRouter {
	controllerMap := make(map[string]IRouter)
	controllerMap["user"] = &user.UserController{}
	return controllerMap
}
