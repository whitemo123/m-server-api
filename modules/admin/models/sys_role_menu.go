package models

type SysRoleMenu struct {
	ID     int64 `json:"id,string" gorm:"primaryKey;"`
	RoleId int64 `json:"roleId,string" gorm:"column:role_id"`
	MenuId int64 `json:"menuId,string" gorm:"column:menu_id"`
}
