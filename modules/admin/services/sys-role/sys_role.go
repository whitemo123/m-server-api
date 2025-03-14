package roleService

import (
	"errors"
	"m-server-api/initializers"
	"m-server-api/modules/admin/dtos"
	roleDto "m-server-api/modules/admin/dtos/sys-role"
	"m-server-api/modules/admin/models"
	roleVo "m-server-api/modules/admin/vos/sys-role"
	"m-server-api/utils/jwt"

	"github.com/jinzhu/copier"
	"github.com/spf13/cast"
)

// 插入角色菜单表
func InsertRoleMenu(roleId int64, menuIds []string) {
	var roleMenus []models.SysRoleMenu
	for _, menuId := range menuIds {
		roleMenus = append(roleMenus, models.SysRoleMenu{
			RoleId: roleId,
			MenuId: cast.ToInt64(menuId),
		})
	}
	initializers.DB.Create(&roleMenus)
}

// 删除角色菜单表
func DeleteRoleMenu(roleId int64) {
	initializers.DB.Where("role_id = ?", roleId).Delete(&models.SysRoleMenu{})
}

// 创建
func Create(d *roleDto.CreateDto, sessionUserInfo jwt.SessionUserInfo) (*models.SysRole, error) {
	var exit models.SysRole
	initializers.DB.Where("role_name = ? AND tenant_id = ?", d.RoleName, sessionUserInfo.TenantId).First(&exit)
	if exit.ID > 0 {
		return nil, errors.New("角色已存在")
	}
	if len(d.MenuIdList) == 0 {
		return nil, errors.New("菜单不能为空")
	}
	role := &models.SysRole{
		RoleName: d.RoleName,
	}
	role.CreateUser = &sessionUserInfo.Id
	role.UpdateUser = &sessionUserInfo.Id
	role.TenantId = &sessionUserInfo.TenantId
	role.Status = d.Status

	result := initializers.DB.Save(role)
	if result.Error != nil {
		return nil, result.Error
	}

	InsertRoleMenu(role.ID, d.MenuIdList)

	return role, nil
}

// 修改
func Modify(d *roleDto.ModifyDto, sessionUserInfo jwt.SessionUserInfo) (*models.SysRole, error) {
	if d.RoleName != "" {
		var exit models.SysRole
		initializers.DB.Where("role_name = ? AND tenant_id = ? AND id != ?", d.RoleName, sessionUserInfo.TenantId, d.ID).First(&exit)
		if exit.ID > 0 {
			return nil, errors.New("角色已存在")
		}
	}
	var role = &models.SysRole{}
	err := initializers.DB.First(role, d.ID).Error
	if err != nil {
		return nil, err
	}
	if d.Status != nil {
		role.Status = d.Status
	}

	if d.RoleName != "" {
		role.RoleName = d.RoleName
	}
	role.UpdateUser = &sessionUserInfo.Id

	if len(d.MenuIdList) > 0 {
		// 删除角色菜单表
		DeleteRoleMenu(role.ID)
		// 插入角色菜单表
		InsertRoleMenu(role.ID, d.MenuIdList)
	}

	result := initializers.DB.Save(role)
	if result.Error != nil {
		return nil, result.Error
	}
	return role, nil
}

// 详情
func Detail(id int64) (*roleVo.SysRoleVo, error) {
	var role models.SysRole
	err := initializers.DB.First(&role, id).Error
	if err != nil {
		return nil, err
	}
	var roleVo roleVo.SysRoleVo
	copier.Copy(&roleVo, &role)

	var roleMenus []models.SysRoleMenu
	err = initializers.DB.Where("role_id = ?", id).Find(&roleMenus).Error
	if err != nil {
		return nil, err
	}
	var menuIds []string
	for _, roleMenu := range roleMenus {
		menuIds = append(menuIds, cast.ToString(roleMenu.MenuId))
	}
	roleVo.MenuIdList = menuIds

	return &roleVo, nil
}

// 删除
func Del(id int64) (bool, error) {
	var role models.SysRole
	err := initializers.DB.Delete(&role, id).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// 列表查询
func List(query roleDto.ListDto, sessionUserInfo jwt.SessionUserInfo) ([]models.SysRole, error) {
	var roles []models.SysRole

	// 查询条件
	db := initializers.DB
	db = db.Where("tenant_id = ?", sessionUserInfo.TenantId)
	if query.Status != nil {
		db = db.Where("status = ?", query.Status)
	}

	db = db.Order("create_time desc")

	err := db.Find(&roles).Error
	if err != nil {
		return []models.SysRole{}, err
	}
	return roles, nil
}

func Page(query roleDto.PageDto, sessionUserInfo jwt.SessionUserInfo) (*dtos.PageRes, error) {
	p, l := dtos.BuildPageQuery(query.Page)
	// 查询条件
	db := initializers.DB

	db = db.Where("tenant_id = ?", sessionUserInfo.TenantId)
	if query.Status != nil {
		db = db.Where("status = ?", query.Status)
	}

	if query.Role != "" {
		db = db.Where("role_name LIKE ? OR id like ?", "%"+query.Role+"%", "%"+query.Role+"%")
	}

	if query.CreateTimeStart != "" {
		db = db.Where("create_time >= ?", query.CreateTimeStart)
	}
	if query.CreateTimeEnd != "" {
		db = db.Where("create_time <= ?", query.CreateTimeEnd)
	}

	db = db.Order("create_time desc")

	offset := (p - 1) * l

	var total int64
	var roles []models.SysRole

	err := db.Offset(offset).Limit(l).Find(&roles).Count(&total).Error
	if err != nil {
		return nil, err
	}

	return &dtos.PageRes{
		Total: total,
		Page:  p,
		Limit: l,
		List:  roles,
	}, nil
}
