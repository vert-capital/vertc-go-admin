package vertc_go_admin

import (
	"gorm.io/gorm"
)

func Migrations(db *gorm.DB) {
	db.AutoMigrate(&Grupo{})
	db.AutoMigrate(&Usuario{})
	db.AutoMigrate(&UsuarioGrupo{})
	db.AutoMigrate(&Table{})
}
