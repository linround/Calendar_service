package middleware

import (
	"calendar_service/model/common/response"
	"calendar_service/utils"
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Token")
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", ctx)
			ctx.Abort()
			return
		}
		j := utils.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), ctx)
			return
		}
		//
		ctx.Set("claims", claims)
	}
}
