package api

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vert-capital/vertc-go-admin/api/handlers"
	"github.com/vert-capital/vertc-go-admin/api/middleware"
	"github.com/vert-capital/vertc-go-admin/infrasctructure/database/repository"
	usecase_usuario "github.com/vert-capital/vertc-go-admin/usecases/usuario"
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

	ucUsuario := usecase_usuario.NewService(
		repository.NewUsuarioPostgres(db),
	)

	group := r.Group("/api/admin/v1")
	group.Use(middleware.AuthMiddleware(ucUsuario))

	handlers.MountTableHandlers(group, db)
	return r
}

func StartWebServer(db *gorm.DB, port string) {
	r := setupRouter(db)
	log.Fatal(r.Run(":" + port))
}
