package system

import (
	"calendar_service/global"
	"calendar_service/model/common/response"
	"calendar_service/model/system/request"
	"github.com/gin-gonic/gin"
)

type CalendarDBApi struct {
}

func (d *CalendarDBApi) InitDB(ctx *gin.Context) {
	if global.CalendarDB != nil {
		response.FailWithMessage("数据库已存在", ctx)
		return
	}
	var dbInfo request.InitDB
	err := ctx.ShouldBindJSON(&dbInfo)
	if err != nil {
		response.FailWithMessage("参数错误", ctx)
		return
	}
	err = calendarDBService.InitDB(dbInfo)
	if err != nil {
		response.FailWithMessage("创建数据库失败", ctx)
		return
	}
	response.OkWithMessage("创建数据库成功", ctx)

}
func (d *CalendarDBApi) CheckDB(ctx *gin.Context) {
	message := "请初始化数据库"
	needInit := true
	if global.CalendarDB != nil {
		message = "数据库已存在"
		needInit = false
	}
	response.OkWithDetailed(gin.H{"needInit": needInit}, message, ctx)
}
