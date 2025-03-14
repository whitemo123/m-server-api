package {{ .Name }}Controller

import (
	adminControllers "m-server-api/modules/admin/controllers"
	{{ .Name }}Dto "m-server-api/modules/admin/dtos/{{ .Folder }}"
	{{ .Name }}Service "m-server-api/modules/admin/services/{{ .Folder }}"
	"m-server-api/utils/excel"
	"m-server-api/utils/resp"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// 分页查询
func Page(c *gin.Context) {
	var query {{ .Name }}Dto.PageDto
	if err := c.ShouldBindQuery(&query); err != nil {
		resp.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	pageRes, err := {{ .Name }}Service.Page(query, adminControllers.GetSessionUserInfo(c))
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Ok(c, pageRes)
}

// 列表查询
func List(c *gin.Context) {
	var query {{ .Name }}Dto.ListDto
	if err := c.ShouldBindQuery(&query); err != nil {
		resp.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	{{ .Name }}s, err := {{ .Name }}Service.List(query, adminControllers.GetSessionUserInfo(c))
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Ok(c, {{ .Name }}s)
}

// 创建
func Create(c *gin.Context) {
	var data {{ .Name }}Dto.CreateDto
	if err := c.ShouldBind(&data); err != nil {
		resp.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	{{ .Name }}, err := {{ .Name }}Service.Create(&data, adminControllers.GetSessionUserInfo(c))
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Ok(c, {{ .Name }})
}

// 编辑
func Modify(c *gin.Context) {
	var data {{ .Name }}Dto.ModifyDto
	if err := c.ShouldBind(&data); err != nil {
		resp.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	{{ .Name }}, err := {{ .Name }}Service.Modify(&data, adminControllers.GetSessionUserInfo(c))
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Ok(c, {{ .Name }})
}

// 详情
func Detail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		resp.Fail(c, http.StatusBadRequest, "id不能为空")
		return
	}
	{{ .Name }}, err := {{ .Name }}Service.Detail(cast.ToInt64(id))
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Ok(c, {{ .Name }})
}

// 删除
func Del(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		resp.Fail(c, http.StatusBadRequest, "id不能为空")
		return
	}
	res, err := {{ .Name }}Service.Del(cast.ToInt64(id))
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Ok(c, res)
}

// 导出
func Export(c *gin.Context) {
	var query {{ .Name }}Dto.ListDto
	if err := c.ShouldBindQuery(&query); err != nil {
		resp.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	{{ .Name }}s, err := {{ .Name }}Service.Export(query, adminControllers.GetSessionUserInfo(c))
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	excel.ExportExcel(c, {{ .Name }}s, "导出")
}
