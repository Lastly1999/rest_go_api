package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "resetgoapi.com/rest_go_api/docs"
	"resetgoapi.com/rest_go_api/middleware"
	"resetgoapi.com/rest_go_api/pkg/config"
	"resetgoapi.com/rest_go_api/routers/routes"
)

func InitRouter() *gin.Engine {
	app := gin.Default()
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1Api := app.Group("/v1")
	authApi := app.Group("/v1")
	authApi.Use(middleware.JwtMiddleware(config.GlobalConfig.Jwt.Isuser))

	routes.RegisterLoginRouter(v1Api)
	routes.RegisterUserRouter(authApi)
	return app
}
