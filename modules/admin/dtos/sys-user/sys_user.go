package userDto

import "m-server-api/modules/admin/dtos"

type CreateDto struct {
	dtos.BaseDto
	Account    string   `json:"account" form:"account" binding:"required"`
	Password   string   `json:"password" form:"password" binding:"required"`
	Name       string   `json:"name" form:"name" binding:"required"`
	Avatar     string   `json:"avatar" form:"avatar"`
	RoleIdList []string `json:"roleIdList" form:"roleIdList" binding:"required"`
}

type ModifyDto struct {
	dtos.BaseDto
	ID         int64    `json:"id,string" form:"id,string" binding:"required"`
	Name       string   `json:"name" form:"name" binding:"required"`
	Avatar     string   `json:"avatar" form:"avatar"`
	RoleIdList []string `json:"roleIdList" form:"roleIdList" binding:"required"`
}

type ListDto struct {
	Status *int `json:"status" form:"status"`
}

type PageDto struct {
	dtos.Page
	Status *int `json:"status" form:"status"`
	// ID/用户名/昵称
	User string `json:"user" form:"user"`
	// 创建时间范围查询
	CreateTimeStart string `json:"createTimeStart" form:"createTimeStart"`
	CreateTimeEnd   string `json:"createTimeEnd" form:"createTimeEnd"`
}
