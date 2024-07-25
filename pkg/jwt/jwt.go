package jwt

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	models "resetgoapi.com/go-admin-api/models/system"
)

type JWTClaims struct {
	User models.SysUser
	jwt.RegisteredClaims
}

func GetClaims(ctx *gin.Context) (*JWTClaims, error) {
	value, exists := ctx.Get("claims")
	if !exists {
		return &JWTClaims{}, errors.New("未找到 claims")
	}
	claims := value.(*JWTClaims)
	return claims, nil
}

// GetUserID 从 Gin 上下文中获取用户 ID
func GetUserID(ctx *gin.Context) (int64, error) {
	claims, err := GetClaims(ctx)
	if err != nil {
		return 0, err
	}

	return claims.User.Id, nil
}
