package routes

import (
	"github.com/gin-gonic/gin"
	v1 "resetgoapi.com/rest_go_api/api/v1/system"
)

func RegisterPostRoutes(route *gin.RouterGroup) {
	postRoute := route.Group("/post")
	postApi := v1.PostApi{}
	{
		postRoute.GET("/page", postApi.Page)
		postRoute.POST("/create", postApi.Create)
		postRoute.PATCH("/update", postApi.Update)
		postRoute.DELETE("/delete", postApi.Delete)
		postRoute.GET("/info/:id", postApi.Info)
	}
}
