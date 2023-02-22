package global

import "time"

type CalendarEventModel struct {
	ID       uint      // 主键
	CreateAt time.Time // 创建时间
	UpdateAt time.Time // 更新时间
	DeleteAt time.Time // 删除时间
}
