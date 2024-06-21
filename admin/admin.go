package vertc_go_admin

import (
	"gorm.io/gorm"
)

var Tables map[string]Table

func AddTable(table Table) {
	Tables[table.Name] = table
}

func RunServer(db *gorm.DB) {
	Migrations(db)
	go StartKafka(db)
}
