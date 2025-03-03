package adminRoutes

import (
	"m-server-api/bootstrap"
	tenantController "m-server-api/modules/admin/controllers/sys-tenant"
)

func init() {
	tenant := bootstrap.New("/admin/tenant")
	tenant.GET("/list", tenantController.List)
	tenant.GET("/page", tenantController.Page)
	tenant.GET("/detail/:id", tenantController.Detail)
	tenant.POST("/create", tenantController.Create)
	tenant.POST("/modify", tenantController.Modify)
}
