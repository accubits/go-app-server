package user

import (
	"logger"

	"github.com/gin-gonic/gin"
)

var log = logger.GetLogger()

func TestLog(c *gin.Context) {
	log.Debug("testing 123")
	c.JSON(200, gin.H{"message": "User founded!", "user": c.Param("id")})
}
