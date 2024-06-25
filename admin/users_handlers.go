package vertc_go_admin

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/cas.v2"
	"gorm.io/gorm"
)

type AuthHandlers struct {
	UcUser IUsecaseUsers
}

func NewAuthHandler(usecaseUsuario IUsecaseUsers) *AuthHandlers {
	return &AuthHandlers{UcUser: usecaseUsuario}
}

// @Summary Login user with CAS Server
// @Description Login with CAS Server
// @Tags Auth
// @Accept  json
// @Produce  json
// method: GET
func (a AuthHandlers) CASAuth(c *gin.Context) {
	frontUrl := os.Getenv("FRONTEND_AUTH_URL")

	if !cas.IsAuthenticated(c.Request) {
		cas.RedirectToLogin(c.Writer, c.Request)
		return
	}

	redirect, existe := c.GetQuery("redirect")
	if existe {
		frontUrl = redirect
	}

	userName := cas.Username(c.Request)

	usuario, err := a.UcUser.GetByEmail(userName)

	if exception := handleError(c, err); exception {
		return
	}

	token, refreshToken, err := JWTTokenGenerator(*usuario)

	if exception := handleError(c, err); exception {
		return
	}

	frontUrlRedirect := fmt.Sprintf("%s/?token=%s&refreshToken=%s", frontUrl, token, refreshToken)

	println("Redirecting to: ", frontUrlRedirect)
	c.Redirect(302, frontUrlRedirect)
}

func (h AuthHandlers) CASLogout(c *gin.Context) {
	c.Set("usuario", nil)
	cas.RedirectToLogout(c.Writer, c.Request)
}

func MountCASAuthRoutes(r *gin.Engine, conn *gorm.DB) {

	usecaseUsuario := NewUserService(
		NewRepositoryUsers(conn),
	)

	authHandlers := NewAuthHandler(usecaseUsuario)

	group := r.Group("/")
	group.Use(SSOAuthMiddleware())

	group.GET("/api/admin/auth/cas", authHandlers.CASAuth)
	group.GET("/api/admin/auth/cas/logout", authHandlers.CASLogout)
}
