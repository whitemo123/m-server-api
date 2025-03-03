package userController

import (
	adminControllers "m-server-api/modules/admin/controllers"
	userDto "m-server-api/modules/admin/dtos/sys-user"
	userService "m-server-api/modules/admin/services/sys-user"
	"m-server-api/utils/excel"
	"m-server-api/utils/resp"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func Create(c *gin.Context) {
	var data userDto.CreateDto
	if err := c.ShouldBind(&data); err != nil {
		resp.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := userService.Create(&data, adminControllers.GetSessionUserInfo(c))
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Ok(c, user)
}

func Modify(c *gin.Context) {
	var data userDto.ModifyDto
	if err := c.ShouldBind(&data); err != nil {
		resp.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := userService.Modify(&data, adminControllers.GetSessionUserInfo(c))
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Ok(c, user)
}

func Detail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		resp.Fail(c, http.StatusBadRequest, "id不能为空")
		return
	}
	user, err := userService.Detail(cast.ToInt64(id))
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Ok(c, user)
}

func Del(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		resp.Fail(c, http.StatusBadRequest, "id不能为空")
		return
	}
	res, err := userService.Del(cast.ToInt64(id))
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Ok(c, res)
}

func List(c *gin.Context) {
	var query userDto.ListDto
	if err := c.ShouldBindQuery(&query); err != nil {
		resp.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	users, err := userService.List(query, adminControllers.GetSessionUserInfo(c))
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Ok(c, users)
}

func Export(c *gin.Context) {
	var query userDto.ListDto
	if err := c.ShouldBindQuery(&query); err != nil {
		resp.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	users, err := userService.Export(query, adminControllers.GetSessionUserInfo(c))
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	excel.ExportExcel(c, users, "用户列表")
}

func Page(c *gin.Context) {
	var query userDto.PageDto
	if err := c.ShouldBindQuery(&query); err != nil {
		resp.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	pageRes, err := userService.Page(query, adminControllers.GetSessionUserInfo(c))
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Ok(c, pageRes)
}
