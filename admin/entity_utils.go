package vertc_go_admin

import "reflect"

func getFields(structModel interface{}) *Fields {
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
		required := false
		if typeOfTable.Field(i).Tag.Get("validate") == "required" {
			required = true
		}
		fieldInfo := map[string]interface{}{
			"type":       fieldType,
			"required":   required,
			"json_field": typeOfTable.Field(i).Tag.Get("json"),
		}

		campos[fieldName] = fieldInfo
	}
	return &campos
}

func addTabela(table Table) {
	if tables == nil {
		tables = make(Tables)
	}
	tables[table.Name] = table
}

func SetAdmin(obj_table interface{}, table_name string, category *string, search_fields []string, actions map[string](func(ids []string) error)) {
	new_table := Table{}
	fields := getFields(obj_table)
	new_table.Fields = fields
	new_table.Name = table_name
	if category != nil {
		new_table.Category = *category
	} else {
		new_table.Category = "ADMIN"
	}
	new_table.SearchFields = search_fields
	new_table.Actions = (*Actions)(&actions)
	addTabela(new_table)
}
