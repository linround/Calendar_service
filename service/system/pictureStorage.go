package system

import (
	"calendar_service/global"
	"calendar_service/model/system/request"
	"calendar_service/system"
)

type PictureStorageService struct {
}

func (service *PictureStorageService) AddPicture(picture system.PictureStorage) (err error) {
	return global.CalendarDB.Create(&picture).Error
}

func (service *PictureStorageService) GetPictureList(params request.SearchPictureListParams) (list []system.PictureStorage, err error) {

	err = global.CalendarDB.Model(&system.PictureStorage{}).Where("user_id = ?", params.UserID).Find(&list).Error
	return
}
