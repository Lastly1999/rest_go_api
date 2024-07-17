package service

import (
	"errors"
	gojwt "github.com/golang-jwt/jwt/v5"
	models "resetgoapi.com/rest_go_api/models/system"
	"resetgoapi.com/rest_go_api/pkg/config"
	"resetgoapi.com/rest_go_api/pkg/jwt"
	"time"
)

type JwtService struct {
}

type IJwtService interface {
	CheckToken(key any, jwtStr string, options ...gojwt.ParserOption) (interface{}, error)
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
