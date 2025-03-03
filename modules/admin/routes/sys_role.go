package adminRoutes

import (
	"m-server-api/bootstrap"
	roleController "m-server-api/modules/admin/controllers/sys-role"
)

func init() {
	role := bootstrap.New("/admin/role")
	role.GET("/list", roleController.List)
	role.GET("/page", roleController.Page)
	role.GET("/detail/:id", roleController.Detail)
	role.GET("/del/:id", roleController.Del)
	role.POST("/create", roleController.Create)
	role.POST("/modify", roleController.Modify)
}
