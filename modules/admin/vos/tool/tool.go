package toolVo

type DataBaseTableVo struct {
	TableName    string `json:"tableName" gorm:"column:table_name"`
	TableComment string `json:"tableComment" gorm:"column:table_comment"`
}

type TableColumnVo struct {
	DataType      string `json:"dataType" gorm:"column:data_type"`
	ColumnName    string `json:"columnName" gorm:"column:column_name"`
	DataTypeLong  string `json:"dataTypeLong" gorm:"column:data_type_long"`
	ColumnComment string `json:"columnComment" gorm:"column:column_comment"`
	PrimaryKey    bool   `json:"primaryKey" gorm:"column:primary_key"`
}
