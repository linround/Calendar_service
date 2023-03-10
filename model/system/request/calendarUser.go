package request

type SearchCalendarUserParams struct {
}

type Register struct {
	Password    string `json:"password"`    //  - 符号代表注释掉 注册时 可更改
	UserAccount string `json:"userAccount"` // 用户账户  注册时  不可更改
	UserName    string `json:"userName"`    // 用户名    可设置 可更改
	AvatarUrl   string `json:"avatarUrl"`   // 头像地址  可设置 可更改
	UserEmail   string `json:"userEmail"`   // 用户邮箱  可设置 可更改
}
