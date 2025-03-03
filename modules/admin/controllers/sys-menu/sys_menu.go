package menuController

import (
	adminControllers "m-server-api/modules/admin/controllers"
	menuDto "m-server-api/modules/admin/dtos/sys-menu"
	menuService "m-server-api/modules/admin/services/sys-menu"
	"m-server-api/utils/resp"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// 菜单树形结构
func Tree(c *gin.Context) {
	menuTree, err := menuService.Tree(adminControllers.GetSessionUserInfo(c))
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Ok(c, menuTree)
}

// 创建菜单
func Create(c *gin.Context) {
	var data menuDto.CreateDto
	if err := c.ShouldBind(&data); err != nil {
		resp.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	menu, err := menuService.Create(&data, adminControllers.GetSessionUserInfo(c))
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Ok(c, menu)
}

// 修改菜单
func Modify(c *gin.Context) {
	var data menuDto.ModifyDto
	if err := c.ShouldBind(&data); err != nil {
		resp.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	menu, err := menuService.Modify(&data, adminControllers.GetSessionUserInfo(c))
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Ok(c, menu)
}

func Detail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		resp.Fail(c, http.StatusBadRequest, "id不能为空")
		return
	}
	menu, err := menuService.Detail(cast.ToInt64(id))
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Ok(c, menu)
}

func Del(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		resp.Fail(c, http.StatusBadRequest, "id不能为空")
		return
	}
	res, err := menuService.Del(cast.ToInt64(id))
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Ok(c, res)
}
