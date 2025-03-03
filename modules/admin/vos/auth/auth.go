package authVo

type LoginVo struct {
	Token string `json:"token"`
}

type UserInfoVo struct {
	ID         int64  `json:"id,string"`
	Account    string `json:"account"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
	TenantId   int64  `json:"tenantId,string"`
	TenantName string `json:"tenantName"`
}
