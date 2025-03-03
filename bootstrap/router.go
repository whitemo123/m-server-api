package bootstrap

import (
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

// 路由信息
type router struct {
	Method       string            //方法名称
	RelativePath string            //url路径
	HandlerFunc  []gin.HandlerFunc //执行函数
}

// 路由组信息
type routerGroup struct {
	RelativePath string            //url路径
	Handlers     []gin.HandlerFunc //中间件
	Router       []*router         //路由信息
}

var GroupList = make([]*routerGroup, 0)

// 创建一个路由组
func New(relativePath string, middleware ...gin.HandlerFunc) *routerGroup {
	var rg routerGroup
	rg.Router = make([]*router, 0)
	rg.RelativePath = relativePath
	rg.Handlers = middleware
	GroupList = append(GroupList, &rg)
	return &rg
}

// 下级分组，继承上级分组的 路径和中间件
func (group *routerGroup) Group(relativePath string, middleware ...gin.HandlerFunc) *routerGroup {
	subGroup := New(group.RelativePath, group.Handlers...)
	subGroup.Handlers = append(subGroup.Handlers, middleware...)
	if strings.HasSuffix(subGroup.RelativePath, "/") {
		if strings.HasPrefix(relativePath, "/") {
			subGroup.RelativePath = subGroup.RelativePath + relativePath[1:]
		} else {
			subGroup.RelativePath = subGroup.RelativePath + relativePath
		}
	} else {
		if strings.HasPrefix(relativePath, "/") {
			subGroup.RelativePath = subGroup.RelativePath + relativePath
		} else {
			subGroup.RelativePath = subGroup.RelativePath + "/" + relativePath
		}
	}

	return subGroup
}

// 添加路由信息
func (group *routerGroup) Handle(method, relativePath string, handlers ...gin.HandlerFunc) *routerGroup {
	var r router
	r.Method = method
	r.RelativePath = relativePath
	r.HandlerFunc = handlers
	group.Router = append(group.Router, &r)
	return group
}

// 添加路由信息-ANY
func (group *routerGroup) ANY(relativePath string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle("ANY", relativePath, handlers...)
	return group
}

// 添加路由信息-GET
func (group *routerGroup) GET(relativePath string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(GET, relativePath, handlers...)
	return group
}

// 添加路由信息-POST
func (group *routerGroup) POST(relativePath string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(POST, relativePath, handlers...)
	return group
}

// 添加路由信息-PUT
func (group *routerGroup) PUT(relativePath string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(PUT, relativePath, handlers...)
	return group
}

// 添加路由信息-DELETE
func (group *routerGroup) DELETE(relativePath string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(DELETE, relativePath, handlers...)
	return group
}
