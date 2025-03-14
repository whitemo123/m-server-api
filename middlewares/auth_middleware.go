package middlewares

import (
	"m-server-api/config"
	"m-server-api/initializers"
	"m-server-api/utils/jwt"
	"m-server-api/utils/log"
	"m-server-api/utils/resp"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
)

// isWhite 是否白名单
func isWhite(path string) bool {
	white := false
	// 白名单列表
	whites := config.Get().Server.White
	for _, v := range whites {
		re, err := regexp.Compile("^/" + config.Get().Server.Prefix + v + "$")
		if err != nil {
			return white
		}
		if re.MatchString(path) {
			white = true
			break
		}
	}
	return white
}

// token 验证中间件
func authTokenMiddleware(c *gin.Context) {
	// 平台类型
	platformName := strings.Split(c.Request.URL.Path, "/")[2]
	// 获取请求头Authorization
	token := c.GetHeader(config.AUTHORIZATION)
	if token == "" {
		resp.Fail(c, http.StatusUnauthorized, "用户未登录")
		c.Abort()
		return
	}
	// 判断是否Bearer开头
	if len(token) < 7 || token[:7] != "Bearer " {
		log.Error("用户凭证格式错误")
		resp.Fail(c, http.StatusUnauthorized, "用户凭证格式错误")
		c.Abort()
		return
	}
	// 解析token
	tokenClaims, err := jwt.ParseToken(token[7:])
	if err != nil {
		log.Error("用户凭证解析失败", err)
		resp.Fail(c, http.StatusUnauthorized, "无效的用户凭证")
		c.Abort()
		return
	}

	userInfo := tokenClaims.SessionUserInfo

	// 验证平台类型
	if userInfo.Platform != platformName {
		log.Error("用户凭证平台类型错误")
		resp.Fail(c, http.StatusUnauthorized, "无效的用户凭证")
		c.Abort()
		return
	}

	// 从redis中获取token进行对比
	var redisToken string
	redisToken, err = initializers.RDB.Get(initializers.Ctx, "admin-token:"+cast.ToString(userInfo.Id)).Result()
	if err != nil {
		if err == redis.Nil {
			log.Error("redis未找到" + cast.ToString(userInfo.Id) + "的token")
			resp.Fail(c, http.StatusUnauthorized, "无效的用户凭证")
			c.Abort()
			return
		}
		log.Error("redis获取token失败：" + err.Error())
		resp.Fail(c, http.StatusUnauthorized, "无效的用户凭证")
		c.Abort()
		return
	}
	if token[7:] != redisToken {
		log.Error("redis存储token与用户凭证不一致")
		resp.Fail(c, http.StatusUnauthorized, "无效的用户凭证")
		c.Abort()
		return
	}

	c.Set(config.CLAIMS, userInfo)
	c.Next()
}

// 权限 验证中间件
func authPermissionMiddleware(c *gin.Context) {
	// TODO: 权限认证
	c.Next()
}

// auth 验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 白名单放行
		if pass := isWhite(c.Request.URL.Path); pass {
			c.Next()
			return
		}
		// token校验
		authTokenMiddleware(c)
		// 权限校验
		authPermissionMiddleware(c)
	}
}
