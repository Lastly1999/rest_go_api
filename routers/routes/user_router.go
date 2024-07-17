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
		userRouter.POST("/create", userApi.Create)
		userRouter.PATCH("/update", userApi.Update)
		userRouter.GET("/info/:id", userApi.Info)
		userRouter.DELETE("/delete/:id", userApi.Delete)
	}
}
