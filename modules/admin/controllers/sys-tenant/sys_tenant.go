package tenantController

import (
	adminControllers "m-server-api/modules/admin/controllers"
	tenantDto "m-server-api/modules/admin/dtos/sys-tenant"
	tenantService "m-server-api/modules/admin/services/sys-tenant"
	"m-server-api/utils/resp"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func Create(c *gin.Context) {
	var data tenantDto.CreateDto
	if err := c.ShouldBind(&data); err != nil {
		resp.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	tenant, err := tenantService.Create(&data, adminControllers.GetSessionUserInfo(c))
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Ok(c, tenant)
}

func Modify(c *gin.Context) {
	var data tenantDto.ModifyDto
	if err := c.ShouldBind(&data); err != nil {
		resp.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	tenant, err := tenantService.Modify(&data, adminControllers.GetSessionUserInfo(c))
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
	tenant, err := tenantService.Detail(cast.ToInt64(id))
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Ok(c, tenant)
}

func List(c *gin.Context) {
	var query tenantDto.ListDto
	if err := c.ShouldBindQuery(&query); err != nil {
		resp.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	tenants, err := tenantService.List(query, adminControllers.GetSessionUserInfo(c))
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Ok(c, tenants)
}

func Page(c *gin.Context) {
	var query tenantDto.PageDto
	if err := c.ShouldBindQuery(&query); err != nil {
		resp.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	pageRes, err := tenantService.Page(query, adminControllers.GetSessionUserInfo(c))
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Ok(c, pageRes)
}
