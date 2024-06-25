package vertc_go_admin

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoutersAdmin(r *gin.Engine, db *gorm.DB) *gin.Engine {
	ucUsuario := NewService(
		NewService(db),
	)

	group := r.Group("/api/admin/v1")
	group.Use(AuthMiddleware(ucUsuario))

	MountTableHandlers(group, db)
	return r
}
