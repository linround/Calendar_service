package utils

import (
	"calendar_service/model/system/request"
	"github.com/gin-gonic/gin"
)

func GetClaims(ctx *gin.Context) (*request.CustomClaims, error) {
	token := ctx.Request.Header.Get("Token")
	j := NewJWT()
	return j.ParseToken(token)
}

func GetUserAccount(ctx *gin.Context) string {
	claims, ok := ctx.Get("claims")
	// 如果没有直接从上下文获取到改值
	// 有可能是某个新的请求
	if !ok {
		cl, err := GetClaims(ctx)
		if err != nil {
			return ""
		}
		return cl.UserAccount
	}
	waitUse := claims.(*request.CustomClaims)
	return waitUse.UserAccount
}
func GetUserID(ctx *gin.Context) uint64 {
	claims, ok := ctx.Get("claims")
	if !ok {
		cl, err := GetClaims(ctx)
		if err != nil {
			return 0
		}
		return cl.UserID
	}
	waitUse := claims.(*request.CustomClaims)
	return waitUse.UserID
}
