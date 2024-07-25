package routes

import (
	"github.com/gin-gonic/gin"
	v1 "resetgoapi.com/go-admin-api/api/v1/system"
)

func RegisterDeptRoutes(route *gin.RouterGroup) {
	deptRoute := route.Group("/dept")
	deptApi := v1.DeptApi{}
	{
		deptRoute.GET("/info/:id", deptApi.Info)
		deptRoute.POST("/create", deptApi.Create)
		deptRoute.PATCH("/update", deptApi.Update)
		deptRoute.GET("/page", deptApi.Page)
		deptRoute.GET("/list", deptApi.List)
		deptRoute.DELETE("/delete/:id", deptApi.Delete)
	}
}
