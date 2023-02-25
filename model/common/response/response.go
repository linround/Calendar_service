package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 200
)

func Result(code int, data interface{}, msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}
func Ok(ctx *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", ctx)
}
func OkWithMessage(message string, ctx *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, ctx)
}
func OkWithData(data interface{}, ctx *gin.Context) {
	Result(SUCCESS, data, "查询成功", ctx)
}
func OkWithDetailed(data interface{}, message string, ctx *gin.Context) {
	Result(SUCCESS, data, message, ctx)
}

func Fail(ctx *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", ctx)
}
func FailWithMessage(message string, ctx *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, ctx)
}
func FailWithDetailed(data interface{}, message string, ctx *gin.Context) {
	Result(ERROR, data, message, ctx)
}
