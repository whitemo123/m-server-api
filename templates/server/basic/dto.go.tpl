package {{ .Name }}Dto

import "m-server-api/modules/admin/dtos"

// 分页查询参数
type PageDto struct {
	dtos.Page
	Status *int `json:"status" form:"status"`
    // ...
}

// 列表查询参数
type ListDto struct {
	Status *int `json:"status" form:"status"`
    // ...
}

// 创建参数
type CreateDto struct {
	dtos.BaseDto
	// ...
}

// 编辑参数
type ModifyDto struct {
	dtos.BaseDto
	// ...
}
