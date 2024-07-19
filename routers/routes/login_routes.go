package routes

import (
	"github.com/gin-gonic/gin"
	v1 "resetgoapi.com/rest_go_api/api/v1/system"
)

func RegisterLoginRoutes(route *gin.RouterGroup) {
	loginRoute := route.Group("/login")
	loginApi := v1.LoginApi{}
	{
		loginRoute.POST("/sign", loginApi.Login)
	}
}
