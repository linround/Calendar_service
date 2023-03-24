package utils

import (
	"calendar_service/global"
	"calendar_service/model/system/request"
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

var (
	TokenInvalid = errors.New("token 不合法")
)

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.CalendarConfig.JWT.SigningKey),
	}
}

// 创建一个账户密码声明

func (j *JWT) CreateClaims(baseClaims request.BaseClaims) (claims request.CustomClaims) {
	claims = request.CustomClaims{
		BaseClaims: baseClaims,
	}
	return
}

func (j *JWT) CreateToken(claims request.CustomClaims) (token string, err error) {
	// HS256 使用同一个密钥进行签名与验证，适合集中式认证；不适合任何分布式架构；
	// RS256 使用RSA私钥进行签名，使用RSA公钥进行验证 ；RS256 可以将验证委托给其他应用，只要将公钥给他们就行；
	// ES256 与RS256一样；
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, _ := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if token != nil {
		claims, ok := token.Claims.(*request.CustomClaims)
		if ok {
			return claims, nil
		}
	}
	return nil, TokenInvalid
}
