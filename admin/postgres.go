package vertc_go_admin

import (
	"gorm.io/gorm"
)

func Migrations(db *gorm.DB) {
	db.AutoMigrate(&UserSSO{})
}
