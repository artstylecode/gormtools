package logic

import (
	"github.com/artstylecode/artcoding-go/utils"
	"github.com/artstylecode/gormtools/common"
	"regexp"
	"strings"
)

type TableLogic struct {
}

type tableInfo struct {
	Field   string `gorm:"column:Field"`
	Type    string `gorm:"column:Type"`
	Null    string `gorm:"column:Null"`
	Key     string `gorm:"column:Key"`
	Default string `gorm:"column:Default"`
	Extra   string `gorm:"column:Extra"`
}

// getTableInfo 获取表格信息
func (t TableLogic) getTableInfo(tablename string) []tableInfo {
	db := utils.GetDb("conf/config.ini")
	rows, err := db.Raw("desc supp_admin").Rows()
	defer rows.Close()
	results := make([]tableInfo, 0)
	for rows.Next() {
		result := tableInfo{}
		db.ScanRows(rows, &result)
		results = append(results, result)
	}
	utils.FailOnError(err, "")

	return results
}

// EntityInfo 生成实体文件的信息
type EntityInfo struct {
	Name   string
	Fields []FieldInfo
}

//FieldInfo 生成实体字段信息
type FieldInfo struct {
	FieldName      string
	TypeStr        string
	GormMappedInfo string
}

// getName 转换下滑线为驼峰
func (t TableLogic) getName(name string) string {
	return common.ToHumpName(name)
}

// getType 获取系统类型
func (t TableLogic) getType(typeInfo string) string {

	TypeMapped := map[string]string{
		"int":             "int",
		"tinyint":         "int",
		"integer":         "int",
		"bigint":          "int",
		"smallint":        "int",
		"mediumint":       "int",
		"varchar":         "string",
		"text":            "int",
		"longtext":        "string",
		"mediumtext":      "string",
		"char":            "string",
		"linestring":      "string",
		"multilinestring": "string",
		"decimal":         "float32",
		"float":           "float32",
		"double":          "float32",
		"real":            "float32",
		"numeric":         "float32",
		"time":            "time.Time",
		"date":            "time.Time",
		"datetime":        "time.Time",
		"timestamp":       "time.Time",
	}
	var typeStr string
	var ok bool
	if typeStr, ok = TypeMapped[typeInfo]; ok {
	} else {
		re := regexp.MustCompile("([A-Za-z]+)\\(")
		str := string(re.Find([]byte(typeInfo)))

		str = strings.Replace(str, "(", "", -1)
		str = strings.TrimSpace(str)
		typeStr, ok = TypeMapped[str]
	}

	if !ok {
		typeStr = "interface{}"
	}
	return typeStr
}

// getGormRelation 获取orm关系映射
func (t TableLogic) getGormRelation(info *tableInfo) string {
	relationStr := "`gorm:\"column:" + info.Field + "\"`"
	return relationStr
}

// GetEntityInfo 获取实体文件生成信息
func (t TableLogic) GetEntityInfo(name string) *EntityInfo {
	tableInfoList := t.getTableInfo(name)
	entityInfo := EntityInfo{}
	entityInfo.Name = name
	for _, info := range tableInfoList {
		fieldInfo := FieldInfo{}
		fieldInfo.GormMappedInfo = t.getGormRelation(&info)
		fieldInfo.FieldName = t.getName(info.Field)
		fieldInfo.TypeStr = t.getType(info.Type)

		entityInfo.Fields = append(entityInfo.Fields, fieldInfo)
	}
	return &entityInfo
}
