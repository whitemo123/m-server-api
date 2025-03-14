package toolService

import (
	"m-server-api/config"
	"m-server-api/initializers"
	toolDto "m-server-api/modules/admin/dtos/tool"
	toolVo "m-server-api/modules/admin/vos/tool"
	"m-server-api/utils/file"
	"os"
	"path/filepath"
	"text/template"
)

// 获取数据表列表
func DataBaseTableList() ([]toolVo.DataBaseTableVo, error) {
	var dataBaseTableList []toolVo.DataBaseTableVo
	err := initializers.DB.Raw(`
		SELECT
			table_name AS table_name,
			table_comment AS table_comment
		FROM
			information_schema.TABLES 
		WHERE
			table_schema = ?
	`, config.Get().Database.MySQL.Name).Scan(&dataBaseTableList).Error
	if err != nil {
		return nil, err
	}
	return dataBaseTableList, nil
}

// 获取数据表列列表
func TableColumnList(name string) ([]toolVo.TableColumnVo, error) {
	var tableColumnList []toolVo.TableColumnVo
	err := initializers.DB.Raw(`
		SELECT 
			c.COLUMN_NAME column_name,
			c.DATA_TYPE data_type,
			CASE c.DATA_TYPE
				WHEN 'longtext' THEN c.CHARACTER_MAXIMUM_LENGTH
				WHEN 'varchar' THEN c.CHARACTER_MAXIMUM_LENGTH
				WHEN 'double' THEN CONCAT_WS(',', c.NUMERIC_PRECISION, c.NUMERIC_SCALE)
				WHEN 'decimal' THEN CONCAT_WS(',', c.NUMERIC_PRECISION, c.NUMERIC_SCALE)
				WHEN 'int' THEN c.NUMERIC_PRECISION
				WHEN 'bigint' THEN c.NUMERIC_PRECISION
				ELSE '' 
			END AS data_type_long,
			c.COLUMN_COMMENT column_comment,
			CASE WHEN kcu.COLUMN_NAME IS NOT NULL THEN 1 ELSE 0 END AS primary_key,
			c.ORDINAL_POSITION
		FROM 
			INFORMATION_SCHEMA.COLUMNS c
		LEFT JOIN 
			INFORMATION_SCHEMA.KEY_COLUMN_USAGE kcu 
		ON 
			c.TABLE_SCHEMA = kcu.TABLE_SCHEMA 
			AND c.TABLE_NAME = kcu.TABLE_NAME 
			AND c.COLUMN_NAME = kcu.COLUMN_NAME 
			AND kcu.CONSTRAINT_NAME = 'PRIMARY'
		WHERE 
			c.TABLE_NAME = ? 
			AND c.TABLE_SCHEMA = ?
		ORDER BY 
			c.ORDINAL_POSITION;
	`, name, config.Get().Database.MySQL.Name).Scan(&tableColumnList).Error
	if err != nil {
		return nil, err
	}
	return tableColumnList, nil
}

// 生成前端基础模板代码
func CreateBasicCode(d *toolDto.CreateDto) error {
	frontWebRoot := d.ProjectPath
	// src目录路径
	srcPath := filepath.Join(frontWebRoot, "src")
	// views目录路径
	viewsPath := filepath.Join(srcPath, "views/"+d.Folder)
	// apis目录路径
	apisPath := filepath.Join(srcPath, "apis/"+d.Folder)
	apiModulePath := filepath.Join(apisPath, d.FileName)
	if !file.FileExists(viewsPath) {
		err := file.CreateDirectory(viewsPath)
		if err != nil {
			return err
		}
	}
	if !file.FileExists(apisPath) {
		err := file.CreateDirectory(apisPath)
		if err != nil {
			return err
		}
	}
	if !file.FileExists(apiModulePath) {
		err := file.CreateDirectory(apiModulePath)
		if err != nil {
			return err
		}
	}

	// ======================= 生成vue文件
	var err error
	var tpl *template.Template
	vueTplPath := filepath.Join("templates/front/basic", "view.vue.tpl")
	tpl, err = template.ParseFiles(vueTplPath)
	if err != nil {
		return err
	}
	vueOutPath := filepath.Join(viewsPath, d.FileName+".vue")
	var outFile *os.File

	outFile, err = os.Create(vueOutPath)
	if err != nil {
		return err
	}
	defer outFile.Close()
	err = tpl.Execute(outFile, d)
	if err != nil {
		return nil
	}
	// ==================================================
	// ======================= 生成api文件
	apiTplPath := filepath.Join("templates/front/basic", "api.ts.tpl")
	tpl, err = template.ParseFiles(apiTplPath)
	if err != nil {
		return err
	}
	apiOutPath := filepath.Join(apiModulePath, "index.ts")
	outFile, err = os.Create(apiOutPath)
	if err != nil {
		return err
	}
	defer outFile.Close()
	err = tpl.Execute(outFile, d)
	if err != nil {
		return nil
	}
	// ==================================================
	// ======================= 生成types文件
	apiTypesTplPath := filepath.Join("templates/front/basic", "types.ts.tpl")
	tpl, err = template.ParseFiles(apiTypesTplPath)
	if err != nil {
		return err
	}
	apiTypesOutPath := filepath.Join(apiModulePath, "types.ts")
	outFile, err = os.Create(apiTypesOutPath)
	if err != nil {
		return err
	}
	defer outFile.Close()
	err = tpl.Execute(outFile, d)
	if err != nil {
		return nil
	}
	// ==================================================
	return nil
}
