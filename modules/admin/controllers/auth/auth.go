package authController

import (
	adminControllers "m-server-api/modules/admin/controllers"
	authDto "m-server-api/modules/admin/dtos/auth"
	authService "m-server-api/modules/admin/services/auth"
	"m-server-api/utils/resp"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 验证码
func Captcha(c *gin.Context) {
	id, base64s, err := authService.Captcha()
	if err != nil {
		resp.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	resp.Ok(c, gin.H{
		"id":      id,
		"captcha": base64s,
	})
}

// 后台管理登录
func Login(c *gin.Context) {
	var loginParams authDto.LoginDto

	if err := c.ShouldBind(&loginParams); err != nil {
		resp.Fail(c, http.StatusBadRequest, err.Error())
		return
	}

	loginRes, err := authService.Login(loginParams)
	if err != nil {
		resp.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	resp.Ok(c, loginRes)
}

// 用户信息
func UserInfo(c *gin.Context) {
	sessionUserInfo := adminControllers.GetSessionUserInfo(c)
	userInfoVo, err := authService.UserInfo(sessionUserInfo.Id)
	if err != nil {
		resp.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	resp.Ok(c, userInfoVo)
}

// 用户菜单树形结构
func MenuTree(c *gin.Context) {
	sessionUserInfo := adminControllers.GetSessionUserInfo(c)
	menuTree, err := authService.MenuTree(sessionUserInfo.Id)
	if err != nil {
		resp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Ok(c, menuTree)
}

func ChangePassword(c *gin.Context) {
	var changePwdData authDto.ChangePasswordDto

	sessionUserInfo := adminControllers.GetSessionUserInfo(c)

	if err := c.ShouldBind(&changePwdData); err != nil {
		resp.Fail(c, http.StatusBadRequest, err.Error())
		return
	}

	err := authService.ChangePassword(changePwdData, sessionUserInfo.Id)
	if err != nil {
		resp.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	resp.Ok(c, true)
}
