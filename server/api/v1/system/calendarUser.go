package system

import (
	"calendar_service/model/common/response"
	"calendar_service/model/system/request"
	"calendar_service/system"
	"github.com/gin-gonic/gin"
)

type CalendarUserApi struct {
}

func (api *CalendarUserApi) RegisterUser(ctx *gin.Context) {
	var calendarUser request.Register
	err := ctx.ShouldBindJSON(&calendarUser)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	user := system.CalendarUser{
		Password:    calendarUser.Password,
		UserAccount: calendarUser.UserAccount,
		UserName:    calendarUser.UserName,
		AvatarUrl:   calendarUser.AvatarUrl,
		UserEmail:   calendarUser.UserEmail,
	}
	err = calendarUserService.RegisterUser(user)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithMessage("创建成功", ctx)
}

func (api *CalendarUserApi) GetUser(ctx *gin.Context) {
	var params request.SearchCalendarUserParams
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	user, err := calendarUserService.GetUser(params)
	if err != nil {
		response.FailWithMessage("获取失败", ctx)
		return
	}
	response.OkWithDetailed(user, "获取成功", ctx)
}

func (api *CalendarUserApi) UpdateUser(ctx *gin.Context) {
	var params system.CalendarUser
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	err = calendarUserService.UpdateUser(params)
	if err != nil {
		response.FailWithMessage("更新失败", ctx)
		return
	}
	response.OkWithMessage("更新成功", ctx)
}

func (api *CalendarUserApi) DeleteUser(ctx *gin.Context) {
	var params system.CalendarUser
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	err = calendarUserService.DeleteUser(params)
	if err != nil {
		response.FailWithMessage("删除失败", ctx)
		return
	}
	response.OkWithMessage("删除成功", ctx)
}
