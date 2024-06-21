package vertc_go_admin

import (
	"gorm.io/gorm"
)

var Tabelas Tables

func RunServer(db *gorm.DB) {
	Migrations(db)
	go StartKafka(db)
}
