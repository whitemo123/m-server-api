package tenantDto

import "m-server-api/modules/admin/dtos"

type CreateDto struct {
	TenantName string `json:"tenantName" form:"tenantName" binding:"required"`
	Status     *int   `json:"status" form:"status"`
}

type ModifyDto struct {
	ID         int64  `json:"id,string" form:"id,string" binding:"required"`
	TenantName string `json:"tenantName" form:"tenantName"`
	Status     *int   `json:"status" form:"status"`
}

type ListDto struct {
	Status *int `json:"status" form:"status"`
}

type PageDto struct {
	dtos.Page
	Status *int `json:"status" form:"status"`
}
