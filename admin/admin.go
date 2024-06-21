package vertc_go_admin

import (
	"gorm.io/gorm"
)

var Tabelas Tables

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
