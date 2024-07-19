package routes

import (
	"github.com/gin-gonic/gin"
	v1 "resetgoapi.com/rest_go_api/api/v1/system"
)

func RegisterMenuRoutes(route *gin.RouterGroup) {
	menuRoute := route.Group("/menu")
	menuApi := v1.MenuApi{}
	{
		menuRoute.GET("/info/:id", menuApi.Info)
		menuRoute.POST("/create", menuApi.Create)
		menuRoute.PATCH("/update", menuApi.Update)
		menuRoute.GET("/page", menuApi.Page)
		menuRoute.DELETE("/delete/:id", menuApi.Delete)
	}
}
