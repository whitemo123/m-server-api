package toolController

import (
	toolDto "m-server-api/modules/admin/dtos/tool"
	toolService "m-server-api/modules/admin/services/tool"
	"m-server-api/utils/resp"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取数据表列表
func DataBaseTableList(c *gin.Context) {
	dataBaseTableList, err := toolService.DataBaseTableList()
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Ok(c, dataBaseTableList)
}

// 获取数据表列列表
func TableColumnList(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		resp.Fail(c, http.StatusBadRequest, "name不能为空")
		return
	}
	tableColumnList, err := toolService.TableColumnList(name)
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Ok(c, tableColumnList)
}

// 生成前端基础代码
func CreateFrontBasicCode(c *gin.Context) {
	var data toolDto.CreateDto
	if err := c.ShouldBind(&data); err != nil {
		resp.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	err := toolService.CreateBasicCode(&data)
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Ok(c, true)
}
