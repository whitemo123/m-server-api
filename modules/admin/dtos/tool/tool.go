package toolDto

// 字段信息
type Columns struct {
	Label    string `json:"label"`
	Prop     string `json:"prop"`
	Type     string `json:"type"`
	Required bool   `json:"required"`
	Default  string `json:"default"`
	Search   bool   `json:"search"`
	Table    bool   `json:"table"`
	Add      bool   `json:"add"`
	Edit     bool   `json:"edit"`
}

// 前端代码生成参数
type CreateDto struct {
	// 小驼峰
	Name string `json:"name"`
	// 大驼峰
	FName string `json:"fName"`
	// 模块名
	Folder string `json:"folder"`
	// 文件名
	FileName string `json:"fileName"`
	// 项目路径
	ProjectPath string `json:"projectPath"`
	// 字段列表
	Columns []Columns `json:"columns"`
}
