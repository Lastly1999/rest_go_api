package routes

import (
	"github.com/gin-gonic/gin"
	v1 "resetgoapi.com/go-admin-api/api/v1/system"
)

func RegisterDictDataRoutes(route *gin.RouterGroup) {
	dictDataRoute := route.Group("/dictData")
	dictDataApi := v1.DictDataApi{}
	{
		dictDataRoute.GET("/info/:id", dictDataApi.Info)
		dictDataRoute.POST("/create", dictDataApi.Create)
		dictDataRoute.PATCH("/update", dictDataApi.Update)
		dictDataRoute.GET("/page", dictDataApi.Page)
		dictDataRoute.DELETE("/delete/:id", dictDataApi.Delete)
	}
}
