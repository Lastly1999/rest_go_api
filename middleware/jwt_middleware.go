package middleware

import (
	"github.com/gin-gonic/gin"
	gojwt "github.com/golang-jwt/jwt/v5"
	"resetgoapi.com/rest_go_api/pkg/config"
	service "resetgoapi.com/rest_go_api/services/system"
	"resetgoapi.com/rest_go_api/utils/result"
	"strings"
)

func JwtMiddleware(guardName string) gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenStr := context.GetHeader(config.GlobalConfig.Jwt.HeaderKey)
		if tokenStr == "" {
			result.HttpResultAuthError(context, config.GlobalConfig.Jwt.HeaderKey+"不存在!")
			return
		}
		tokenSplits := strings.Split(tokenStr, "Bearer ")
		if len(tokenSplits) != 2 {
			result.HttpResultAuthError(context, config.GlobalConfig.Jwt.HeaderKey+"参数格式有误!")
			return
		}
		token := tokenSplits[1]
		jwtService := service.JwtService{}
		Claims, err := jwtService.CheckToken([]byte(config.GlobalConfig.Jwt.SecretKey), token)
		if err != nil {
			result.HttpResultAuthError(context, "登录已超时，请重新登录！")
			return
		}
		claims := Claims.(gojwt.MapClaims)
		isuser, _ := claims.GetIssuer()
		if guardName != isuser {
			result.HttpResultAuthError(context, "签发isuser存在异常，请重试！")
			return
		}
		context.Set("claims", claims)
	}
}
