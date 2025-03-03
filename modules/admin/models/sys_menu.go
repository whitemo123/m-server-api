package models

type SysMenu struct {
	BaseModel
	ParentId *int64 `json:"parentId,string" gorm:"parent_id"`
	Name     string `json:"name" gorm:"name"`
	Sort     int    `json:"sort" gorm:"sort"`
	Type     int    `json:"type" gorm:"type"`
	Icon     string `json:"icon" gorm:"icon"`
	Path     string `json:"path" gorm:"path"`
	Alias    string `json:"alias" gorm:"alias"`
}
