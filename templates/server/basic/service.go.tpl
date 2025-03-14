package {{ .Name }}Service

import (
	"errors"
	"m-server-api/initializers"
	"m-server-api/modules/admin/dtos"
	{{ .Name }}Dto "m-server-api/modules/admin/dtos/{{ .Folder }}"
	"m-server-api/modules/admin/models"
	{{ .Name }}Vo "m-server-api/modules/admin/vos/{{ .Folder }}"
	md5Encrypt "m-server-api/utils/encrypt/md5"
	"m-server-api/utils/jwt"

	"github.com/jinzhu/copier"
	"github.com/spf13/cast"
)

// 分页查询
func Page(query {{ .Name }}Dto.PageDto, sessionUserInfo jwt.SessionUserInfo) (*dtos.PageRes, error) {
	p, l := dtos.BuildPageQuery(query.Page)
	// 查询条件
	db := initializers.DB
	db = db.Where("tenant_id = ?", sessionUserInfo.TenantId)
	if query.Status != nil {
		db = db.Where("status = ?", query.Status)
	}
    // ...

	offset := (p - 1) * l
	var total int64
	var users []models.{{ .TFName }}
    // 排序
	db = db.Order("create_time desc")

    // 分页查询
	err := db.Offset(offset).Limit(l).Find(&{{ .Name }}s).Count(&total).Error
	if err != nil {
		return nil, err
	}

    // 转换数据
	var res []{{ .Name }}Vo.{{ .FName }}Vo
	copier.Copy(&res, &{{ .Name }}s)

	return &dtos.PageRes{
		Total: total,
		Page:  p,
		Limit: l,
		List:  res,
	}, nil
}

// 列表查询
func List(query {{ .Name }}Dto.ListDto, sessionUserInfo jwt.SessionUserInfo) ([]{{ .Name }}Vo.{{ .FName }}Vo, error) {
	var {{ .Name }}s []models.{{ .FName }}

	// 查询条件
	db := initializers.DB
	db = db.Where("tenant_id = ?", sessionUserInfo.TenantId)
	if query.Status != nil {
		db = db.Where("status = ?", query.Status)
	}
    // ...

    // 排序
	db = db.Order("create_time desc")

    // 查询
	err := db.Find(&{{ .Name }}s).Error
	if err != nil {
		return nil, err
	}
    
    // 转换数据
	var res []{{ .Name }}Vo.{{ .FName }}Vo
	copier.Copy(&res, &{{ .Name }}s)

	return res, nil
}

// 创建
func Create(d *{{ .Name }}Dto.CreateDto, sessionUserInfo jwt.SessionUserInfo) (*models.{{ .TFName }}, error) {
	{{ .Name }} := &models.{{ .TFName }}{
		// ...
	}
    // 通用字段
	{{ .Name }}.TenantId = &sessionUserInfo.TenantId
	{{ .Name }}.CreateUser = &sessionUserInfo.Id
	{{ .Name }}.UpdateUser = &sessionUserInfo.Id
	{{ .Name }}.Status = d.Status

    // 创建数据
	result := initializers.DB.Save({{ .Name }})
	if result.Error != nil {
		return nil, result.Error
	}

	return {{ .Name }}, nil
}

// 修改
func Modify(d *{{ .Name }}Dto.ModifyDto, sessionUserInfo jwt.SessionUserInfo) (*models.{{ .TFName }}, error) {
	var {{ .Name }} = &models.{{ .TFName }}{}
    // 查询出数据
	err := initializers.DB.First({{ .Name }}, d.ID).Error
	if err != nil {
		return nil, err
	}
    // 通用字段
	{{ .Name }}.UpdateUser = &sessionUserInfo.Id
	if d.Status != nil {
		{{ .Name }}.Status = d.Status
	}

	// ...

    // 更新数据
	result := initializers.DB.Save({{ .Name }})
	if result.Error != nil {
		return nil, result.Error
	}

	return {{ .Name }}, nil
}

// 详情
func Detail(id int64) (*{{ .Name }}Vo.{{ .FName }}Vo, error) {
	var {{ .Name }} models.{{ .TFName }}
    // 查询详情
	err := initializers.DB.First(&{{ .Name }}, id).Error
	if err != nil {
		return nil, err
	}
    // 转换数据
	var result {{ .Name }}Vo.{{ .FName }}Vo
	copier.Copy(&result, &{{ .Name }})

	return &result, nil
}

// 删除
func Del(id int64) (bool, error) {
	var {{ .Name }} models.{{ .TFName }}
    // 删除数据
	err := initializers.DB.Delete(&{{ .Name }}, id).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

// 导出excel
func Export(query {{ .Name }}Dto.ListDto, sessionUserInfo jwt.SessionUserInfo) ([][]interface{}, error) {
	var {{ .Name }}s []models.{{ .TFName }}

	// 查询条件
	db := initializers.DB
	db = db.Where("tenant_id = ?", sessionUserInfo.TenantId)
	if query.Status != nil {
		db = db.Where("status = ?", query.Status)
	}
    // ...

    // 排序
	db = db.Order("create_time desc")

    // 查询数据
	err := db.Find(&{{ .Name }}s).Error
	if err != nil {
		return nil, err
	}

	var res [][]interface{}
    // 表头
	res = append(res, []interface{}{
        // ...
    })
    // 添加数据
	for _, {{ .Name }} := range {{ .Name }}s {
		res = append(res, []interface{}{
			// ...
		})
	}
	return res, nil
}
