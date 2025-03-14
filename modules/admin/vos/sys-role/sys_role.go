package roleVo

import "m-server-api/modules/admin/models"

type SysRoleVo struct {
	models.SysRole
	MenuIdList []string `json:"menuIdList" form:"menuIdList"`
}
