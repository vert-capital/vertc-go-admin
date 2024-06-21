package main

import (
	"github.com/vert-capital/vertc-go-admin/api"
	"github.com/vert-capital/vertc-go-admin/entity"
	postgres "github.com/vert-capital/vertc-go-admin/infrasctructure/database"
	"github.com/vert-capital/vertc-go-admin/infrasctructure/kafka"
	"gorm.io/gorm"
)

var Tables map[string]entity.Table

var DB *gorm.DB

func AddTable(table entity.Table) {
	Tables[table.Name] = table
}

func SetDB(db *gorm.DB) {
	DB = db
}

func main() {
}

func RunServer(db *gorm.DB, port string) {
	postgres.Migrations(db)
	go kafka.StartKafka(db)
	api.StartWebServer(db, port)
}
