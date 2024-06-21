package middleware

import (
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	usecase_usuario "github.com/vert-capital/vertc-go-admin/usecases/usuario"

	"gopkg.in/cas.v2"
)

func AuthMiddleware(usecase usecase_usuario.IUsecaseUsuario) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get bearer token from header
		bearerToken := c.Request.Header.Get("Authorization")

		usuario, err := usecase.GetUsuarioByToken(bearerToken)

		// check if token is valid
		if err == nil && usuario.IsAdmin {

			// set usuario to context
			c.Set("usuario", *usuario)

			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
		}
	}
}

func SSOAuthMiddleware() gin.HandlerFunc {
	casURL := os.Getenv("CAS_SERVER_URL")

	url, _ := url.Parse(casURL)

	client := cas.NewClient(&cas.Options{
		URL: url,
	})
	handler := client.HandleFunc(func(writer http.ResponseWriter, request *http.Request) {
		// DO NOTHING
	})

	return func(ctx *gin.Context) {
		// Call the normal method
		handler.ServeHTTP(ctx.Writer, ctx.Request)

		if !cas.IsAuthenticated(ctx.Request) {
			// redirect to login url
			cas.RedirectToLogin(ctx.Writer, ctx.Request)
			return
		}
		ctx.Next()
	}

}
