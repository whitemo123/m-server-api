package userService

import (
	"errors"
	"m-server-api/initializers"
	"m-server-api/modules/admin/dtos"
	userDto "m-server-api/modules/admin/dtos/sys-user"
	"m-server-api/modules/admin/models"
	md5Encrypt "m-server-api/utils/encrypt/md5"
	"m-server-api/utils/jwt"

	"github.com/spf13/cast"
)

// 插入用户角色表
func InsertUserRole(userId int64, roleIds []string) {
	var userRoles []models.SysUserRole
	for _, roleId := range roleIds {
		userRoles = append(userRoles, models.SysUserRole{
			UserId: userId,
			RoleId: cast.ToInt64(roleId),
		})
	}
	initializers.DB.Create(&userRoles)
}

// 删除用户角色表
func DeleteUserRole(userId int64) {
	initializers.DB.Where("user_id = ?", userId).Delete(&models.SysUserRole{})
}

// 创建
func Create(d *userDto.CreateDto, sessionUserInfo jwt.SessionUserInfo) (*models.SysUser, error) {
	var exit models.SysUser
	initializers.DB.Where("account = ? AND tenant_id = ?", d.Account, sessionUserInfo.TenantId).First(&exit)
	if exit.ID > 0 {
		return nil, errors.New("用户名已存在")
	}

	if len(d.RoleIdList) == 0 {
		return nil, errors.New("角色不能为空")
	}

	// 密码加密
	password := md5Encrypt.Encode(d.Password)

	user := &models.SysUser{
		Account:  d.Account,
		Password: password,
		Name:     d.Name,
		Avatar:   d.Avatar,
	}
	user.TenantId = &sessionUserInfo.TenantId
	user.CreateUser = &sessionUserInfo.Id
	user.UpdateUser = &sessionUserInfo.Id
	user.Status = d.Status

	result := initializers.DB.Save(user)
	if result.Error != nil {
		return nil, result.Error
	}

	// 插入用户角色表
	InsertUserRole(user.ID, d.RoleIdList)

	return user, nil
}

// 修改
func Modify(d *userDto.ModifyDto, sessionUserInfo jwt.SessionUserInfo) (*models.SysUser, error) {
	var user = &models.SysUser{}
	err := initializers.DB.First(user, d.ID).Error
	if err != nil {
		return nil, err
	}
	user.UpdateUser = &sessionUserInfo.Id
	if d.Status != nil {
		user.Status = d.Status
	}
	if d.Name != "" {
		user.Name = d.Name
	}
	if d.Avatar != "" {
		user.Avatar = d.Avatar
	}

	if len(d.RoleIdList) > 0 {
		// 删除用户角色表
		DeleteUserRole(user.ID)
		// 插入用户角色表
		InsertUserRole(user.ID, d.RoleIdList)
	}

	result := initializers.DB.Save(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

// 详情
func Detail(id int64) (*models.SysUser, error) {
	var user models.SysUser
	err := initializers.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// 删除
func Del(id int64) (bool, error) {
	var user models.SysUser
	err := initializers.DB.Delete(&user, id).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// 列表查询
func List(query userDto.ListDto, sessionUserInfo jwt.SessionUserInfo) ([]models.SysUser, error) {
	var users []models.SysUser

	// 查询条件
	db := initializers.DB
	db = db.Where("tenant_id = ?", sessionUserInfo.TenantId)
	if query.Status != nil {
		db = db.Where("status = ?", query.Status)
	}

	db = db.Order("create_time desc")

	err := db.Find(&users).Error
	if err != nil {
		return []models.SysUser{}, err
	}
	return users, nil
}

// 导出excel
func Export(query userDto.ListDto, sessionUserInfo jwt.SessionUserInfo) ([][]interface{}, error) {
	var users []models.SysUser

	// 查询条件
	db := initializers.DB
	db = db.Where("tenant_id = ?", sessionUserInfo.TenantId)
	if query.Status != nil {
		db = db.Where("status = ?", query.Status)
	}

	db = db.Order("create_time desc")

	err := db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	var res [][]interface{}
	res = append(res, []interface{}{"ID", "账号", "姓名"})
	for _, user := range users {
		res = append(res, []interface{}{
			cast.ToString(user.ID),
			user.Account,
			user.Name,
		})
	}
	return res, nil
}

// 分页查询
func Page(query userDto.PageDto, sessionUserInfo jwt.SessionUserInfo) (*dtos.PageRes, error) {
	p, l := dtos.BuildPageQuery(query.Page)
	// 查询条件
	db := initializers.DB
	db = db.Where("tenant_id = ?", sessionUserInfo.TenantId)

	if query.Status != nil {
		db = db.Where("status = ?", query.Status)
	}

	offset := (p - 1) * l

	var total int64
	var users []models.SysUser

	db = db.Order("create_time desc")

	err := db.Offset(offset).Limit(l).Find(&users).Count(&total).Error
	if err != nil {
		return nil, err
	}

	return &dtos.PageRes{
		Total: total,
		Page:  p,
		Limit: l,
		List:  users,
	}, nil
}
