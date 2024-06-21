package vertc_go_admin

import (
	api "github.com/vert-capital/vertc-go-admin/api"
	entity "github.com/vert-capital/vertc-go-admin/entity"
	postgres "github.com/vert-capital/vertc-go-admin/infrasctructure/database"
	kafka "github.com/vert-capital/vertc-go-admin/infrasctructure/kafka"
	"gorm.io/gorm"
)

var Tables map[string]entity.Table

func AddTable(table entity.Table) {
	Tables[table.Name] = table
}

func RunServer(db *gorm.DB, port string) {
	postgres.Migrations(db)
	go kafka.StartKafka(db)
	api.StartWebServer(db, port)
}
