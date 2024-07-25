package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "resetgoapi.com/go-admin-api/docs"
	"resetgoapi.com/go-admin-api/middleware"
	"resetgoapi.com/go-admin-api/pkg/config"
	"resetgoapi.com/go-admin-api/routers/routes"
)

func InitRouter() *gin.Engine {
	app := gin.Default()
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1Api := app.Group("/v1")
	authApi := app.Group("/v1")
	authApi.Use(middleware.JwtMiddleware(config.GlobalConfig.Jwt.Isuser))

	routes.RegisterLoginRoutes(v1Api)
	routes.RegisterUserRoutes(authApi)
	routes.RegisterPostRoutes(authApi)
	routes.RegisterRoleRoutes(authApi)
	routes.RegisterMenuRoutes(authApi)
	routes.RegisterDeptRoutes(authApi)
	routes.RegisterDictDataRoutes(authApi)
	routes.RegisterDictTypeRoutes(authApi)
	return app
}
