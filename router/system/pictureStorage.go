package system

import (
	v1 "calendar_service/server/api/v1"
	"github.com/gin-gonic/gin"
)

type PictureStorageRouter struct {
}

func (receiver *PictureStorageRouter) InitPictureStorageRouter(Router *gin.RouterGroup) {
	pictureStorageRouter := Router.Group("picture")
	pictureStorageApi := v1.ApiGroupApp.PictureStorageApi
	pictureStorageRouter.POST("add", pictureStorageApi.AddPicture)
	pictureStorageRouter.GET("list", pictureStorageApi.GetPictureList)
}
