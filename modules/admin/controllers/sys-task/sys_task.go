package taskController

import (
	taskService "m-server-api/modules/admin/services/task"
	"m-server-api/utils/resp"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func List(c *gin.Context) {
	tasks := taskService.GetTaskList()
	resp.Ok(c, tasks)
}

func SwitchTaskStatus(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		resp.Fail(c, 400, "id不能为空")
		return
	}
	taskService.SwitchTaskStatus(cast.ToInt64(id))
	resp.Ok(c, true)
}
