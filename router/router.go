package router

import (
	"github.com/gin-gonic/gin"
	v1 "resetgoapi.com/rest_go_api/api/v1"
)

func InitRouter() *gin.Engine {
	app := gin.Default()
	app.GET("/test", v1.Test)
	return app
}
