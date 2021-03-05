package main

import (
	"github.com/artstylecode/gormtools/logic"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Result struct {
	Field   string `gorm:"column:Field"`
	Type    string `gorm:"column:Type"`
	Null    string `gorm:"column:Null"`
	Key     string `gorm:"column:Key"`
	Default string `gorm:"column:Default"`
	Extra   string `gorm:"column:Extra"`
}

func main() {
	generatorLogic := logic.GeneratorLogic{}
	generatorLogic.CreateEntityFile("./entity/model.go", "supp_admin", "entity")

}
