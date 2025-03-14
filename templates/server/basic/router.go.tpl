package adminRoutes

import (
	"m-server-api/bootstrap"
	{{ .Name }}Controller "m-server-api/modules/admin/controllers/{{ .Folder }}"
)

func init() {
	{{ .Name }} := bootstrap.New("/admin/{{ .Name }}")
	{{ .Name }}.GET("/list", {{ .Name }}Controller.List)
	{{ .Name }}.GET("/export", {{ .Name }}Controller.Export)
	{{ .Name }}.GET("/page", {{ .Name }}Controller.Page)
	{{ .Name }}.GET("/detail/:id", {{ .Name }}Controller.Detail)
	{{ .Name }}.GET("/del/:id", {{ .Name }}Controller.Del)
	{{ .Name }}.POST("/create", {{ .Name }}Controller.Create)
	{{ .Name }}.POST("/modify", {{ .Name }}Controller.Modify)
}
