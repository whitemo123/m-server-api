package adminRoutes

import (
	"m-server-api/bootstrap"
	userController "m-server-api/modules/admin/controllers/sys-user"
)

func init() {
	user := bootstrap.New("/admin/user")
	user.GET("/list", userController.List)
	user.GET("/export", userController.Export)
	user.GET("/page", userController.Page)
	user.GET("/detail/:id", userController.Detail)
	user.POST("/create", userController.Create)
	user.POST("/modify", userController.Modify)
}
