package global

import (
	"gorm.io/gorm"
	"time"
)

type CalendarEventModel struct {
	ID        uint           `gorm:"primarykey"` // 主键
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index"` // 删除时间
}
