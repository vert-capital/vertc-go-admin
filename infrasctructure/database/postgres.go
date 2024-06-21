package infrascructure

import (
	"github.com/vert-capital/vertc-go-admin/entity"
	"gorm.io/gorm"
)

func Migrations(db *gorm.DB) {
	db.AutoMigrate(&entity.Grupo{})
	db.AutoMigrate(&entity.Usuario{})
	db.AutoMigrate(&entity.UsuarioGrupo{})
}
