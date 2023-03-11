package global

import (
	"gorm.io/gorm"
	"time"
)

type CalendarEventModel struct {
	// json 代表的是转成json后的显示字段
	CreatedAt time.Time      `json:"-"`              // 创建时间
	UpdatedAt time.Time      `json:"-"`              // 更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"` // 删除时间
}
