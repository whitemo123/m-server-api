package adminControllers

import (
	"m-server-api/config"
	"m-server-api/utils/jwt"

	"github.com/gin-gonic/gin"
)

// 获取上下文用户信息
func GetSessionUserInfo(c *gin.Context) jwt.SessionUserInfo {
	userInfo, _ := c.Get(config.CLAIMS)
	return userInfo.(jwt.SessionUserInfo)
}
