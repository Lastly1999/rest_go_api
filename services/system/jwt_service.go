package service

import (
	"errors"
	gojwt "github.com/golang-jwt/jwt/v5"
	models "resetgoapi.com/go-admin-api/models/system"
	"resetgoapi.com/go-admin-api/pkg/config"
	"resetgoapi.com/go-admin-api/pkg/jwt"
	"time"
)

var JwtService = jwtService{}

type jwtService struct {
}

type IJwtService interface {
	CheckToken(key any, jwtStr string) (*jwt.JWTClaims, error)
	GetToken(user *models.SysUser) string
}

// CheckToken godoc
// 验证token
func (s *jwtService) CheckToken(key any, jwtStr string) (*jwt.JWTClaims, error) {
	token, err := gojwt.ParseWithClaims(jwtStr, &jwt.JWTClaims{}, func(token *gojwt.Token) (interface{}, error) {
		return key, nil
	})
	claims, ok := token.Claims.(*jwt.JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, err
}

// GetToken godoc
// 派发token
func (s *jwtService) GetToken(user *models.SysUser) string {
	claims := jwt.JWTClaims{
		User: *user,
		RegisteredClaims: gojwt.RegisteredClaims{
			ExpiresAt: gojwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(config.GlobalConfig.Jwt.ExpiresIn))), // 过期时间
			IssuedAt:  gojwt.NewNumericDate(time.Now()),                                                                     // 签发时间
			NotBefore: gojwt.NewNumericDate(time.Now()),                                                                     // 生效时间
			Issuer:    config.GlobalConfig.Jwt.Isuser,                                                                       // 签发人
		},
	}
	token := gojwt.NewWithClaims(gojwt.SigningMethodHS256, claims)
	signToken, err := token.SignedString([]byte(config.GlobalConfig.Jwt.SecretKey))
	if err != nil {
		return ""
	}
	return signToken
}
