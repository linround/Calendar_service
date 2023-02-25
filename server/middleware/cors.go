package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 配置跨域请求和options方法

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		origin := ctx.Request.Header.Get("Origin")
		// 跨域配置
		ctx.Header("Access-Control-Allow-Origin", origin)
		//	 options方法配置
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		ctx.Next()
	}
}
