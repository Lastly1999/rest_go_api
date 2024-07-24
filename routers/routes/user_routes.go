package routes

import (
	"github.com/gin-gonic/gin"
	v1 "resetgoapi.com/rest_go_api/api/v1/system"
)

func RegisterUserRoutes(route *gin.RouterGroup) {
	userRoute := route.Group("/user")
	userApi := v1.UserApi{}
	{
		userRoute.GET("/page", userApi.Page)
		userRoute.POST("/create", userApi.Create)
		userRoute.PATCH("/update", userApi.Update)
		userRoute.GET("/info/:id", userApi.Info)
		userRoute.DELETE("/delete/:id", userApi.Delete)
		userRoute.GET("/getUserInfo", userApi.GetUserInfo)
	}
}
