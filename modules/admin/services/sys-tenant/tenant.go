package tenantService

import (
	"errors"
	"m-server-api/initializers"
	"m-server-api/modules/admin/dtos"
	tenantDto "m-server-api/modules/admin/dtos/sys-tenant"
	"m-server-api/modules/admin/models"
	"m-server-api/utils/jwt"
)

// 创建
func Create(d *tenantDto.CreateDto, sessionUserInfo jwt.SessionUserInfo) (*models.SysTenant, error) {
	var exitTenant models.SysTenant
	initializers.DB.Where("tenant_name = ?", d.TenantName).First(&exitTenant)
	if exitTenant.ID > 0 {
		return nil, errors.New("租户已存在")
	}
	tenant := &models.SysTenant{
		TenantName: d.TenantName,
	}
	tenant.Status = d.Status
	tenant.CreateUser = &sessionUserInfo.Id
	tenant.UpdateUser = &sessionUserInfo.Id

	result := initializers.DB.Model(&models.SysTenant{}).Create(tenant)
	if result.Error != nil {
		return nil, result.Error
	}
	return tenant, nil
}

// 修改
func Modify(d *tenantDto.ModifyDto, sessionUserInfo jwt.SessionUserInfo) (*models.SysTenant, error) {
	if d.TenantName != "" {
		var exitTenant models.SysTenant
		initializers.DB.Where("tenant_name = ? AND id != ?", d.TenantName, d.ID).First(&exitTenant)
		if exitTenant.ID > 0 {
			return nil, errors.New("租户已存在")
		}
	}

	var tenant models.SysTenant
	err := initializers.DB.First(&tenant, d.ID).Error
	if err != nil {
		return nil, err
	}
	tenant.UpdateUser = &sessionUserInfo.Id
	if d.Status != nil {
		tenant.Status = d.Status
	}
	if d.TenantName != "" {
		tenant.TenantName = d.TenantName
	}

	result := initializers.DB.Save(&tenant)
	if result.Error != nil {
		return nil, result.Error
	}
	return &tenant, nil
}

// 详情
func Detail(id int64) (*models.SysTenant, error) {
	var tenant models.SysTenant
	err := initializers.DB.First(&tenant, id).Error
	if err != nil {
		return nil, err
	}
	return &tenant, nil
}

// 删除
func Del(id int64) (bool, error) {
	var tenant models.SysTenant
	err := initializers.DB.Delete(&tenant, id).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// 列表查询
func List(query tenantDto.ListDto, sessionUserInfo jwt.SessionUserInfo) ([]models.SysTenant, error) {
	var tenants []models.SysTenant

	// 查询条件
	db := initializers.DB
	if query.Status != nil {
		db = db.Where("status = ?", query.Status)
	}

	db = db.Order("create_time desc")

	err := db.Find(&tenants).Error
	if err != nil {
		return []models.SysTenant{}, err
	}
	return tenants, nil
}

func Page(query tenantDto.PageDto, sessionUserInfo jwt.SessionUserInfo) (*dtos.PageRes, error) {
	p, l := dtos.BuildPageQuery(query.Page)
	// 查询条件
	db := initializers.DB

	if query.Status != nil {
		db = db.Where("status = ?", query.Status)
	}

	db = db.Order("create_time desc")

	offset := (p - 1) * l

	var total int64
	var tenants []models.SysTenant

	err := db.Offset(offset).Limit(l).Find(&tenants).Count(&total).Error
	if err != nil {
		return nil, err
	}

	return &dtos.PageRes{
		Total: total,
		Page:  p,
		Limit: l,
		List:  tenants,
	}, nil
}
