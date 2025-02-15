package middleware

import (
	"github.com/gin-gonic/gin"
	"resetgoapi.com/go-admin-api/pkg/config"
	service "resetgoapi.com/go-admin-api/services/system"
	"resetgoapi.com/go-admin-api/utils/result"
	"strings"
)

func JwtMiddleware(guardName string) gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenStr := context.GetHeader(config.GlobalConfig.Jwt.HeaderKey)
		jsonResult := result.JsonResult{Context: context}
		if tokenStr == "" {
			jsonResult.HttpResultAuthError(config.GlobalConfig.Jwt.HeaderKey + "不存在!")
			return
		}
		tokenSplits := strings.Split(tokenStr, "Bearer ")
		if len(tokenSplits) != 2 {
			jsonResult.HttpResultAuthError(config.GlobalConfig.Jwt.HeaderKey + "参数格式有误!")
			return
		}
		token := tokenSplits[1]
		Claims, err := service.JwtService.CheckToken([]byte(config.GlobalConfig.Jwt.SecretKey), token)
		if err != nil {
			jsonResult.HttpResultAuthError("登录已超时，请重新登录！")
			return
		}
		if guardName != Claims.Issuer {
			jsonResult.HttpResultAuthError("签发isuser存在异常，请重试！")
			return
		}
		context.Set("claims", Claims)
	}
}
