package roleDto

import "m-server-api/modules/admin/dtos"

type CreateDto struct {
	dtos.BaseDto
	RoleName   string   `json:"roleName" form:"roleName" binding:"required"`
	MenuIdList []string `json:"menuIdList" form:"menuIdList" binding:"required"`
}

type ModifyDto struct {
	dtos.BaseDto
	RoleName   string   `json:"roleName" form:"roleName"`
	ID         int64    `json:"id,string" form:"id,string" binding:"required"`
	MenuIdList []string `json:"menuIdList" form:"menuIdList" binding:"required"`
}

type ListDto struct {
	Status *int `json:"status" form:"status"`
}

type PageDto struct {
	dtos.Page
	Status *int `json:"status" form:"status"`
}
