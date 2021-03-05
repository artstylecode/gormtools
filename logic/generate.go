package logic

import (
	"fmt"
	"github.com/artstylecode/artcoding-go/file"
)

// GeneratorLogic 生成文件逻辑
type GeneratorLogic struct {
}

// CreateEntityFile 创建实体文件
func (t GeneratorLogic) CreateEntityFile(output string, tableName string, packageName string) {
	tableLogic := TableLogic{}
	entityInfo := tableLogic.GetEntityInfo(tableName)
	entityContent := fmt.Sprintf("package %s\r\ntype %s struct {\r\n", packageName, entityInfo.Name)
	for _, fieldInfo := range entityInfo.Fields {
		fieldItemStr := fmt.Sprintf("%s %s %s", fieldInfo.FieldName, fieldInfo.TypeStr, fieldInfo.GormMappedInfo)
		entityContent = fmt.Sprintf("%s %s\r\n", entityContent, fieldItemStr)
	}
	entityContent += "}"
	txtFile := file.TextFile{}
	txtFile.SaveFile(output, entityContent)

}
