package models

type SysTenant struct {
	TenantName string `json:"tenantName" gorm:"column:tenant_name"`
	BaseModelNoTenant
}
