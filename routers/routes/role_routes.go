package routes

import (
	"github.com/gin-gonic/gin"
	v1 "resetgoapi.com/rest_go_api/api/v1/system"
)

func RegisterRoleRoutes(route *gin.RouterGroup) {
	roleRoute := route.Group("/role")
	roleApi := v1.RoleApi{}
	{
		roleRoute.GET("/page", roleApi.Page)
		roleRoute.POST("/create", roleApi.Create)
		roleRoute.PATCH("/update", roleApi.Update)
		roleRoute.DELETE("/delete/:id", roleApi.Delete)
		roleRoute.GET("/info/:id", roleApi.Info)
	}
}
