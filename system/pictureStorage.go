package system

import "calendar_service/global"

type PictureStorage struct {
	global.CalendarEventModel `yaml:",inline"`
	PictureId                 uint64 `json:"pictureId" gorm:"not null;unique;primary_key;comment:记录ID;size:90"`
	UserId                    uint64 `json:"userId" gorm:"图片所属的用户ID"`
	ImageUrl                  string `json:"imageUrl" gorm:"图片链接"`
	PictureType               string `json:"pictureType" gorm:"图片的类型"`
	Prompt                    string `json:"prompt" gorm:"图片的生成提示词"`
	SourceIP                  string `json:"sourceIP" gorm:"访问请求的IP"`
}

type PictureStorageRecords struct {
	Total int64 `json:"total" gorm:"统计总条数"`
}

func (s PictureStorage) TableName() string {
	return "picture_storage"
}
