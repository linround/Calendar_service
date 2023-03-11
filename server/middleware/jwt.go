package middleware

import (
	"calendar_service/model/common/response"
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("X-TOKEN")
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", ctx)
			ctx.Abort()
			return
		}
	}
}
