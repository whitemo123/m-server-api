package adminRoutes

import (
	"m-server-api/bootstrap"
	toolController "m-server-api/modules/admin/controllers/tool"
)

func init() {
	tool := bootstrap.New("/admin/tool")
	tool.GET("/dataBaseTableList", toolController.DataBaseTableList)
	tool.GET("/tableColumnList", toolController.TableColumnList)
	tool.POST("/createFrontBasicCode", toolController.CreateFrontBasicCode)
}
