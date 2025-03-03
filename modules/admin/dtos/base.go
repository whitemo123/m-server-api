package dtos

type Page struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}

type PageRes struct {
	Total int64       `json:"total"`
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
	List  interface{} `json:"list"`
}

type BaseDto struct {
	Status *int `json:"status" form:"status"`
}

// 构建分页查询
func BuildPageQuery(pageQuery Page) (page int, limit int) {
	var p int
	var l int
	if pageQuery.Page <= 0 {
		p = 1
	} else {
		p = pageQuery.Page
	}
	if pageQuery.Limit <= 0 {
		l = 10
	} else {
		l = pageQuery.Limit
	}
	return p, l
}
