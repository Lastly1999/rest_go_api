package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	models "resetgoapi.com/rest_go_api/models/system"
)

type JWTClaims struct {
	User models.SysUser
	jwt.RegisteredClaims
}
