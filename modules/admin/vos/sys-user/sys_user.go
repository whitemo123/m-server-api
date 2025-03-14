package userVo

import "m-server-api/modules/admin/models"

type SysUserVo struct {
	Account    string   `json:"account" gorm:"column:account"`
	Password   string   `json:"-" copier:"-"`
	Name       string   `json:"name" gorm:"column:name"`
	Avatar     string   `json:"avatar" gorm:"column:avatar"`
	RoleIdList []string `json:"roleIdList" form:"roleIdList"`
	models.BaseModel
}
