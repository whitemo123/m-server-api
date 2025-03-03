package models

type SysUser struct {
	Account  string `json:"account" gorm:"column:account"`
	Password string `json:"password" gorm:"column:password"`
	Name     string `json:"name" gorm:"column:name"`
	Avatar   string `json:"avatar" gorm:"column:avatar"`
	BaseModel
}
