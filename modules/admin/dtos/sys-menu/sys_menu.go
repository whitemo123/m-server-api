package menuDto

import "m-server-api/modules/admin/dtos"

type CreateDto struct {
	dtos.BaseDto
	ParentId *int64 `json:"parentId,string" form:"parentId,string"`
	Name     string `json:"name" form:"name" binding:"required"`
	Sort     int    `json:"sort,string" form:"sort" binding:"required"`
	Type     int    `json:"type" form:"type" binding:"required"`
	Icon     string `json:"icon" form:"icon"`
	Path     string `json:"path" form:"path"`
	Alias    string `json:"alias" form:"alias"`
	Keep     *int   `json:"keep" form:"keep"`
}

type ModifyDto struct {
	dtos.BaseDto
	ParentId *int64 `json:"parentId,string" form:"parentId"`
	ID       int64  `json:"id,string" form:"id,string" binding:"required"`
	Name     string `json:"name" form:"name"`
	Sort     *int   `json:"sort" form:"sort"`
	Icon     string `json:"icon" form:"icon"`
	Path     string `json:"path" form:"path"`
	Alias    string `json:"alias" form:"alias"`
	Keep     *int   `json:"keep" form:"keep"`
}
