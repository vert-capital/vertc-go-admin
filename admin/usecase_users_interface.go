package vertc_go_admin

type IRepositoryUsers interface {
	GetByEmail(email string) (user *UserSSO, err error)
	CreateOrUpdateUser(user *UserSSO) error
}

type IUsecaseUsers interface {
	GetUserByToken(token string) (user *UserSSO, err error)
	CreateOrUpdateUser(user *UserSSO) error
	GetByEmail(email string) (user *UserSSO, err error)
}
