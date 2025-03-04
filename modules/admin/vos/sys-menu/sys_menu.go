package menuVo

type MenuTree struct {
	ID       int64       `json:"id,string"`
	ParentId *int64      `json:"parentId,string"`
	Name     string      `json:"name"`
	Sort     int         `json:"sort"`
	Type     int         `json:"type"`
	Icon     string      `json:"icon"`
	Path     string      `json:"path"`
	Alias    string      `json:"alias"`
	Status   *int        `json:"status"`
	Keep     *int        `json:"keep"`
	Children []*MenuTree `json:"children"`
}
