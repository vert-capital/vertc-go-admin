package vertc_go_admin

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type UseCaseUsers struct {
	repo IRepositoryUsers
}

type LoginDetails struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func NewService(repository IRepositoryUsers) *UseCaseUsers {
	return &UseCaseUsers{repository}
}

func ValidateToken(tokenString string) (claims *LoginDetails, err error) {

	token, err := jwt.ParseWithClaims(
		tokenString,
		&LoginDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY_GESTAO")), nil
		},
	)

	if err, ok := err.(*jwt.ValidationError); ok {
		if err.Errors&jwt.ValidationErrorExpired != 0 {
			return token.Claims.(*LoginDetails), errors.New("token expirado")
		} else {
			return nil, errors.New("token inválido")
		}
	}
	claims, ok := token.Claims.(*LoginDetails)
	if !ok {
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, err
	}

	return claims, nil

}

func (s *UseCaseUsers) GetUserByToken(token string) (user *UserSSO, err error) {
	claims, err := ValidateToken(token)
	if err != nil {
		return nil, err
	}
	user, err = s.repo.GetByEmail(claims.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UseCaseUsers) CreateOrUpdateUser(user *UserSSO) error {
	return s.repo.CreateOrUpdateUser(user)
}

func (s *UseCaseUsers) GetByEmail(email string) (user *UserSSO, err error) {
	return s.repo.GetByEmail(email)
}

func JWTTokenGenerator(u UserSSO) (signedToken string, signedRefreshToken string, err error) {
	claims := LoginDetails{
		Email: *u.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7 * 365).Unix(),
		},
	}

	refreshClaims := LoginDetails{
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
