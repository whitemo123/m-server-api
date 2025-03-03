package models

type SysUserRole struct {
	ID     int64 `json:"id,string" gorm:"primaryKey;"`
	UserId int64 `json:"userId,string" gorm:"column:user_id"`
	RoleId int64 `json:"roleId,string" gorm:"column:role_id"`
}
