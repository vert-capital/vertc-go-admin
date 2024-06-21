package vertc_go_admin

import (
	"gorm.io/gorm"
)

func RunServer(db *gorm.DB) {
	Migrations(db)
	go StartKafka(db)
}
