package service

import (
	"errors"
	gojwt "github.com/golang-jwt/jwt/v5"
	models "resetgoapi.com/rest_go_api/models/system"
	"resetgoapi.com/rest_go_api/pkg/config"
	"resetgoapi.com/rest_go_api/pkg/jwt"
)

type JwtService struct {
}

type IJwtService interface {
	CheckToken(token string) (interface{}, error)
	GetToken(user *models.SysUser) string
}

// CheckToken godoc
// 验证token
func (s *JwtService) CheckToken(key any, jwtStr string, options ...gojwt.ParserOption) (interface{}, error) {
	mc := gojwt.MapClaims{}
	token, err := gojwt.ParseWithClaims(jwtStr, mc, func(token *gojwt.Token) (interface{}, error) {
		return key, nil
	}, options...)
	if err != nil {
		return nil, err
	}
	// 校验 Claims 对象是否有效，基于 exp（过期时间），nbf（不早于），iat（签发时间）等进行判断（如果有这些声明的话）。
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token.Claims, nil
}

// GetToken godoc
// 派发token
func (s *JwtService) GetToken(user *models.SysUser) string {
	claims := jwt.JWTClaims{
		User:             *user,
		RegisteredClaims: gojwt.RegisteredClaims{},
	}
	token := gojwt.NewWithClaims(gojwt.SigningMethodHS256, claims)
	signToken, err := token.SignedString(config.GlobalConfig.Jwt.SecretKey)
	if err != nil {
		return ""
	}
	return signToken
}
