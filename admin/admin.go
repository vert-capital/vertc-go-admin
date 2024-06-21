package vertc_go_admin

import (
	"gorm.io/gorm"
)

var Tables map[string]Table

func AddTable(table Table) {
	Tables[table.Name] = table
}

func RunServer(db *gorm.DB, port string) {
	Migrations(db)
	go StartKafka(db)
}
