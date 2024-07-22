package routes

import (
	"github.com/gin-gonic/gin"
	v1 "resetgoapi.com/rest_go_api/api/v1/system"
)

func RegisterDictTypeRoutes(route *gin.RouterGroup) {
	dictTypeRoute := route.Group("/dictType")
	dictTypeApi := v1.DictDataApi{}
	{
		dictTypeRoute.GET("/info/:id", dictTypeApi.Info)
		dictTypeRoute.POST("/create", dictTypeApi.Create)
		dictTypeRoute.PATCH("/update", dictTypeApi.Update)
		dictTypeRoute.GET("/page", dictTypeApi.Page)
		dictTypeRoute.DELETE("/delete/:id", dictTypeApi.Delete)
	}
}
