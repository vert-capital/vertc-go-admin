package vertc_go_admin

import (
	"reflect"

	"gorm.io/gorm"
)

var Tabelas Tables

func GetFields(structModel interface{}) *Fields {
	campos := make(Fields)
	example := structModel
	v := reflect.ValueOf(example)
	typeOfTable := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fieldName := typeOfTable.Field(i).Name

		// Determinando o tipo do campo
		fieldType := typeOfTable.Field(i).Type.Kind().String()

		// Estrutura a ser armazenada no mapa
		if fieldType == "ptr" {
			fieldType = "datetime"
		}
		fieldInfo := map[string]interface{}{
			"type":       fieldType,
			"required":   true,
			"json_field": typeOfTable.Field(i).Tag.Get("json"),
		}

		campos[fieldName] = fieldInfo
	}
	return &campos
}

func AddTabela(tabela Table) {
	if Tabelas == nil {
		Tabelas = make(Tables)
	}
	Tabelas[tabela.Name] = tabela
}

func RunServer(db *gorm.DB) {
	Migrations(db)
	go StartKafka(db)
}
