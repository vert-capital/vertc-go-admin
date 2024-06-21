package vertc_go_admin

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.New()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("authorization")

	r.Use(cors.New(config))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	ucUsuario := NewService(
		NewUsuarioPostgres(db),
	)

	group := r.Group("/api/admin/v1")
	group.Use(AuthMiddleware(ucUsuario))

	MountTableHandlers(group, db)
	return r
}

func StartWebServer(db *gorm.DB, port string) {
	r := setupRouter(db)
	log.Fatal(r.Run(":" + port))
}
