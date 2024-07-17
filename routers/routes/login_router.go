package routes

import (
	"github.com/gin-gonic/gin"
	v1 "resetgoapi.com/rest_go_api/api/v1/system"
)

func RegisterLoginRouter(router *gin.RouterGroup) {
	loginRouter := router.Group("/login")
	loginApi := v1.LoginApi{}
	{
		loginRouter.POST("/sign", loginApi.Login)
	}
}
