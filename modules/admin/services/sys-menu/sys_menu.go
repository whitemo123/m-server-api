package menuService

import (
	"m-server-api/initializers"
	menuDto "m-server-api/modules/admin/dtos/sys-menu"
	"m-server-api/modules/admin/models"
	menuVo "m-server-api/modules/admin/vos/sys-menu"
	"m-server-api/utils/jwt"
	"sort"
)

// 排序树形菜单
func sortMenuTree(nodes []*menuVo.MenuTree) {
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Sort < nodes[j].Sort
	})
	for _, node := range nodes {
		if len(node.Children) > 0 {
			sortMenuTree(node.Children)
		}
	}
}

// 全部树形菜单
func Tree(sessionUserInfo jwt.SessionUserInfo) ([]*menuVo.MenuTree, error) {
	var menuList []models.SysMenu
	err := initializers.DB.Where("tenant_id = ?", sessionUserInfo.TenantId).Find(&menuList).Error
	if err != nil {
		return nil, err
	}
	menuMap := make(map[int64]*menuVo.MenuTree)
	var roots []*menuVo.MenuTree
	for _, m := range menuList {
		node := &menuVo.MenuTree{
			ID:       m.ID,
			Name:     m.Name,
			Sort:     m.Sort,
			Type:     m.Type,
			Icon:     m.Icon,
			Path:     m.Path,
			Alias:    m.Alias,
			Status:   m.Status,
			ParentId: m.ParentId,
			Children: []*menuVo.MenuTree{},
		}
		menuMap[m.ID] = node
	}
	// 构建树形结构
	for _, node := range menuMap {
		if node.ParentId == nil {
			// 一级菜单
			roots = append(roots, node)
		} else {
			// 找到父级，将当前节点挂载到父级的 Children 中
			if parent, ok := menuMap[*node.ParentId]; ok {
				parent.Children = append(parent.Children, node)
			}
		}
	}
	sortMenuTree(roots)
	return roots, nil
}

// 创建
func Create(d *menuDto.CreateDto, sessionUserInfo jwt.SessionUserInfo) (*models.SysMenu, error) {
	menu := &models.SysMenu{
		Name:  d.Name,
		Sort:  d.Sort,
		Type:  d.Type,
		Icon:  d.Icon,
		Path:  d.Path,
		Alias: d.Alias,
	}
	menu.TenantId = &sessionUserInfo.TenantId
	menu.CreateUser = &sessionUserInfo.Id
	menu.UpdateUser = &sessionUserInfo.Id
	menu.Status = d.Status
	menu.ParentId = d.ParentId

	err := initializers.DB.Save(menu).Error
	if err != nil {
		return nil, err
	}
	return menu, nil
}

// 编辑
func Modify(d *menuDto.ModifyDto, sessionUserInfo jwt.SessionUserInfo) (*models.SysMenu, error) {
	var menu models.SysMenu
	err := initializers.DB.First(&menu, d.ID).Error
	if err != nil {
		return nil, err
	}
	if d.Status != nil {
		menu.Status = d.Status
	}
	if d.Name != "" {
		menu.Name = d.Name
	}
	if d.Sort != nil {
		menu.Sort = *d.Sort
	}
	if d.Icon != "" {
		menu.Icon = d.Icon
	}
	if d.Path != "" {
		menu.Path = d.Path
	}
	if d.Alias != "" {
		menu.Alias = d.Alias
	}
	menu.UpdateUser = &sessionUserInfo.Id

	err = initializers.DB.Save(&menu).Error
	if err != nil {
		return nil, err
	}
	return &menu, nil
}

// 详情
func Detail(id int64) (*models.SysMenu, error) {
	var menu models.SysMenu
	err := initializers.DB.First(&menu, id).Error
	if err != nil {
		return nil, err
	}
	return &menu, nil
}

// 删除
func Del(id int64) (bool, error) {
	var menu models.SysMenu
	err := initializers.DB.Delete(&menu, id).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
