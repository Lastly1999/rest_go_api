package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	app := gin.Default()
	return app
}
