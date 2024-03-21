package system

import (
	"calendar_service/model/common/response"
	"calendar_service/model/system/request"
	"calendar_service/system"
	"github.com/gin-gonic/gin"
)

type PictureStorageApi struct {
}

func (receiver *PictureStorageApi) AddPicture(ctx *gin.Context) {
	var picture system.PictureStorage
	err := ctx.ShouldBindJSON(&picture)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	err = pictureStorageService.AddPicture(picture)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithMessage("创建记录", ctx)
}

func (receiver *PictureStorageApi) GetPictureList(ctx *gin.Context) {
	var params request.SearchPictureListParams
	err := ctx.ShouldBindJSON(&params)
	params.UserID = 1 // 暂时采用固定ID的 方式

	list, err := pictureStorageService.GetPictureList(params)
	if err != nil {
		response.FailWithMessage("获取失败", ctx)
		return
	}
	response.OkWithDetailed(response.ListResult{List: list}, "获取成功", ctx)
}
