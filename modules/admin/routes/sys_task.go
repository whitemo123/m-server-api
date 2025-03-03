package adminRoutes

import (
	"m-server-api/bootstrap"
	taskController "m-server-api/modules/admin/controllers/sys-task"
	_ "m-server-api/modules/admin/services/task"
)

func init() {
	task := bootstrap.New("/admin/task")
	task.GET("/list", taskController.List)
	task.GET("/switch/:id", taskController.SwitchTaskStatus)
}
