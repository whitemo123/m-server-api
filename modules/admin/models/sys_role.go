package models

type SysRole struct {
	BaseModel
	RoleName string `json:"roleName" gorm:"column:role_name"`
}
