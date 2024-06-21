package vertc_go_admin

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	entity "github.com/vert-capital/vertc-go-admin/entity"
	"gorm.io/gorm"
)

type UseCaseUsuario struct {
	repo IRepositoryUsuario
}

type DetalhesLogin struct {
	ID    int
	Name  string
	Email string
	jwt.StandardClaims
}

func NewService(repository IRepositoryUsuario) *UseCaseUsuario {
	return &UseCaseUsuario{repo: repository}
}

func (u *UseCaseUsuario) GetUsuarioByToken(token string) (*entity.Usuario, error) {
	claims, err := ValidateToken(token)

	if err != nil {
		return nil, err
	}

	usuario, err := u.repo.GetByEMail(claims.Email)

	if err != nil {
		return nil, err
	}

	return usuario, nil
}

func ValidateToken(tokenString string) (claims *DetalhesLogin, err error) {

	token, err := jwt.ParseWithClaims(
		tokenString,
		&DetalhesLogin{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY_GESTAO")), nil
		},
	)

	if err, ok := err.(*jwt.ValidationError); ok {
		if err.Errors&jwt.ValidationErrorExpired != 0 {
			return token.Claims.(*DetalhesLogin), errors.New("token expirado")
		} else {
			return nil, errors.New("token inválido")
		}
	}
	claims, ok := token.Claims.(*DetalhesLogin)
	if !ok {
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, err
	}

	return claims, nil

}

func (u *UseCaseUsuario) Create(usuario *entity.Usuario) error {
	err := usuario.Validate()

	if err != nil {
		return err
	}

	_, err = u.repo.GetByEMail(usuario.Email)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return u.repo.CreateUsuario(usuario)
	}

	return err
}

func (u *UseCaseUsuario) Update(usuario *entity.Usuario) error {
	err := usuario.Validate()

	if err != nil {
		return err
	}

	return u.repo.UpdateUsuario(usuario)
}

func (u *UseCaseUsuario) Delete(usuario *entity.Usuario) error {
	return u.repo.DeleteUsuario(usuario)
}

func (u *UseCaseUsuario) GetUsuarioByEMail(email string) (usuario *entity.Usuario, err error) {
	return u.repo.GetByEMail(email)
}

func (u *UseCaseUsuario) GetUsuarioByID(id int) (usuario *entity.Usuario, err error) {
	return u.repo.GetByID(id)
}

func (u *UseCaseUsuario) GetPatrimoniosByUsuarioEmail(email string) (patrimonios []uint, err error) {
	return u.repo.GetPatrimoniosByEmail(email)
}

func (u *UseCaseUsuario) CreateOrUpdateUsuario(usuario *entity.Usuario) error {
	return u.repo.CreateOrUpdate(usuario)
}

func (u *UseCaseUsuario) UpdateUsuarioByEmail(usuario *entity.TipoUsuarioKafka) error {
	return u.repo.UpdateByEmail(usuario)
}

func JWTTokenGenerator(u entity.Usuario) (signedToken string, signedRefreshToken string, err error) {
	claims := DetalhesLogin{
		ID:    int(u.ID),
		Name:  u.PrimeiroNome + " " + u.UltimoNome,
		Email: u.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7 * 365).Unix(),
		},
	}

	refreshClaims := DetalhesLogin{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7 * 365).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("SECRET_KEY_GESTAO")))

	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(os.Getenv("SECRET_KEY_GESTAO")))

	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}
