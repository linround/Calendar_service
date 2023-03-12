package system

import (
	"calendar_service/global"
	"calendar_service/model/common/response"
	"calendar_service/model/system/request"
	sysResponse "calendar_service/model/system/response"
	"calendar_service/system"
	"calendar_service/utils"
	"github.com/gin-gonic/gin"
)

type CalendarBaseApi struct {
}

func (base *CalendarBaseApi) Login(ctx *gin.Context) {
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
	base.TokenAssign(ctx, *user)
}

func (base *CalendarBaseApi) TokenAssign(ctx *gin.Context, user system.CalendarUser) {
	j := utils.JWT{
		SigningKey: []byte(global.CalendarConfig.JWT.SigningKey),
	}
	// 声明的相关参数
	claims := j.CreateClaims(request.BaseClaims{
		UserID:      user.UserID,
		UserAccount: user.UserAccount,
	})
	// 根据声明创建token
	token, err := j.CreateToken(claims)
	if err != nil {
		response.FailWithMessage("获取token失败", ctx)
		return
	}
	response.OkWithDetailed(sysResponse.LoginResponse{
		User:  user,
		Token: token,
	}, "登录成功", ctx)
}
