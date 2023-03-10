package request

type Login struct {
	UserAccount string `json:"userAccount"` // 账户名
	Password    string `json:"password"`    // 密码
	captchaId   string `json:"captchaId"`   // 验证码
	Captcha     string `json:"captcha"`     //  验证码ID
}
