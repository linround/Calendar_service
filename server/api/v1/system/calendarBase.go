package system

import (
	"calendar_service/model/common/response"
	"calendar_service/model/system/request"
	"calendar_service/system"
	"github.com/gin-gonic/gin"
)

type CalendarBaseApi struct {
}

func (b *CalendarBaseApi) Login(ctx *gin.Context) {
	var params request.Login
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	u := &system.CalendarUser{Password: params.Password, UserAccount: params.UserAccount}
	user, err := calendarBaseService.Login(u)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithDetailed(user, "获取成功", ctx)
}
