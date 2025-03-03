package roleController

import (
	adminControllers "m-server-api/modules/admin/controllers"
	roleDto "m-server-api/modules/admin/dtos/sys-role"
	roleService "m-server-api/modules/admin/services/sys-role"
	"m-server-api/utils/resp"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func Create(c *gin.Context) {
	var data roleDto.CreateDto
	if err := c.ShouldBind(&data); err != nil {
		resp.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	tenant, err := roleService.Create(&data, adminControllers.GetSessionUserInfo(c))
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Ok(c, tenant)
}

func Modify(c *gin.Context) {
	var data roleDto.ModifyDto
	if err := c.ShouldBind(&data); err != nil {
		resp.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	tenant, err := roleService.Modify(&data, adminControllers.GetSessionUserInfo(c))
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Ok(c, tenant)
}

func Detail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		resp.Fail(c, http.StatusBadRequest, "id不能为空")
		return
	}
	tenant, err := roleService.Detail(cast.ToInt64(id))
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Ok(c, tenant)
}

func Del(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		resp.Fail(c, http.StatusBadRequest, "id不能为空")
		return
	}
	res, err := roleService.Del(cast.ToInt64(id))
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Ok(c, res)
}

func List(c *gin.Context) {
	var query roleDto.ListDto
	if err := c.ShouldBindQuery(&query); err != nil {
		resp.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	tenants, err := roleService.List(query, adminControllers.GetSessionUserInfo(c))
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Ok(c, tenants)
}

func Page(c *gin.Context) {
	var query roleDto.PageDto
	if err := c.ShouldBindQuery(&query); err != nil {
		resp.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	pageRes, err := roleService.Page(query, adminControllers.GetSessionUserInfo(c))
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Ok(c, pageRes)
}
