package routes

import (
	"github.com/gin-gonic/gin"
	v1 "resetgoapi.com/rest_go_api/api/v1/system"
)

func RegisterUserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("/user")
	userApi := v1.UserApi{}
	{
		userRouter.GET("/page", userApi.Page)
		userRouter.GET("/create", userApi.Create)
		userRouter.GET("/update", userApi.Update)
		userRouter.GET("/delete", userApi.Delete)
	}
}
