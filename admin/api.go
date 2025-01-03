package vertc_go_admin

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoutersAdmin(r *gin.Engine, db *gorm.DB, middleware gin.HandlerFunc) *gin.Engine {
	group := r.Group("/api/admin/v1")
	group.Use(middleware)
	MountTableHandlers(group, db)
	return r
}
