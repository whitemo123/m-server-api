package models

import (
	"database/sql/driver"
	"fmt"
	"m-server-api/utils/snowflake"
	"time"

	"gorm.io/gorm"
)

// 自定义时间格式
type LocalTime time.Time

// 转换时间格式
func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

// Value 实现 driver.Valuer 接口
func (t LocalTime) Value() (driver.Value, error) {
	return time.Time(t), nil
}

type BaseModelNoTenant struct {
	ID         int64     `json:"id,string" gorm:"primaryKey;"`
	Status     *int      `json:"status" gorm:"column:status;default:1"`
	CreateUser *int64    `json:"createUser,string" gorm:"column:create_user"`
	UpdateUser *int64    `json:"updateUser,string" gorm:"column:update_user"`
	CreateTime LocalTime `json:"createTime" gorm:"column:create_time"`
	UpdateTime LocalTime `json:"updateTime" gorm:"column:update_time"`
}

type BaseModel struct {
	ID         int64     `json:"id,string" gorm:"primaryKey;"`
	TenantId   *int64    `json:"tenantId,string" gorm:"column:tenant_id;default:88888888"`
	Status     *int      `json:"status" gorm:"column:status;default:1"`
	CreateUser *int64    `json:"createUser,string" gorm:"column:create_user"`
	UpdateUser *int64    `json:"updateUser,string" gorm:"column:update_user"`
	CreateTime LocalTime `json:"createTime" gorm:"column:create_time"`
	UpdateTime LocalTime `json:"updateTime" gorm:"column:update_time"`
}

// gorm创建前 hook
func (m *BaseModel) BeforeCreate(scope *gorm.DB) error {
	// 雪花ID生成
	m.ID = snowflake.GenSnowflakeId()
	// 创建时间
	m.CreateTime = LocalTime(time.Now())
	// 更新时间
	m.UpdateTime = LocalTime(time.Now())
	return nil
}

// gorm更新前 hook
func (m *BaseModel) BeforeUpdate(scope *gorm.DB) error {
	// 更新时间
	m.UpdateTime = LocalTime(time.Now())
	return nil
}
