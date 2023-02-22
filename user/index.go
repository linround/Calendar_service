package user

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Name     string
	Age      int
	Birthday time.Time
}

func CreateUser(db *gorm.DB) {
	user := User{Name: "linyuan", Age: 18, Birthday: time.Now()}
	result := db.Create(&user)
	if result.Error != nil {
		fmt.Println("出现错误")
	}
}
