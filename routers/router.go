package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"resetgoapi.com/rest_go_api/routers/routes"
)

func InitRouter() *gin.Engine {
	app := gin.Default()
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1Api := app.Group("/v1")
	routes.RegisterUserRouter(v1Api)
	return app
}
