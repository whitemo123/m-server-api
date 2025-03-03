package adminRoutes

import (
	"m-server-api/bootstrap"
	authController "m-server-api/modules/admin/controllers/auth"
)

func init() {
	auth := bootstrap.New("/admin/auth")
	auth.GET("/captcha", authController.Captcha)
	auth.POST("/login", authController.Login)
	auth.GET("/userInfo", authController.UserInfo)
	auth.GET("/menuTree", authController.MenuTree)
	auth.POST("changePwd", authController.ChangePassword)
}
