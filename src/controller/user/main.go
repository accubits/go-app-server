package user

import (
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (uc *UserController) GetRoutes(appGroup *gin.RouterGroup) {
	appGroup.GET("/:id", TestLog)
}

func (mc *UserController) InitSubRoutes(appGroup *gin.RouterGroup) {

}
