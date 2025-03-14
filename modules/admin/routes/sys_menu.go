package adminRoutes

import (
	"m-server-api/bootstrap"
	menuController "m-server-api/modules/admin/controllers/sys-menu"
)

func init() {
	menu := bootstrap.New("/admin/menu")
	menu.GET("/tree", menuController.Tree)
	menu.POST("/create", menuController.Create)
	menu.POST("/modify", menuController.Modify)
	menu.GET("/detail/:id", menuController.Detail)
	menu.GET("/del/:id", menuController.Del)
	menu.GET("/roleMenuTree", menuController.RoleMenuTree)
}
