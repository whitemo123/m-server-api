package authDto

type LoginDto struct {
	Account   string `json:"account" form:"account" binding:"required"`
	Password  string `json:"password" form:"password" binding:"required"`
	Captcha   string `json:"captcha" form:"captcha" binding:"required"`
	CaptchaId string `json:"captchaId" form:"captchaId" binding:"required"`
}

type ChangePasswordDto struct {
	Password    string `json:"password" form:"password" binding:"required"`
	NewPassword string `json:"newPassword" form:"newPassword" binding:"required"`
}
