package vertc_go_admin

//go:generate mockgen -destination=../../mocks/mock_usecase_repository_usuario.go -package=mocks app/usecase/usuario IRepositoryUsuario
type IRepositoryUsuario interface {
	GetByID(id int) (usuario *Usuario, err error)
	GetByEMail(email string) (usuario *Usuario, err error)
	CreateUsuario(usuario *Usuario) error
	UpdateUsuario(usuario *Usuario) error
	UpdateByEmail(usuario *TipoUsuarioKafka) error
	DeleteUsuario(usuario *Usuario) error
	GetUsuario(id int) (usuario *Usuario, err error)
	GetUsuarioGrupos(id int) (grupos []Grupo, err error)
	GetPatrimoniosByEmail(email string) (patrimonios []uint, err error)
	CreateOrUpdate(usuario *Usuario) error
	Existe(usuario *Usuario) (bool, error)
}

//go:generate mockgen -destination=../../mocks/mock_usecase_usuario.go -package=mocks app/usecase/usuario IUsecaseUsuario
type IUsecaseUsuario interface {
	GetUsuarioByToken(token string) (*Usuario, error)
	Create(usuario *Usuario) error
	CreateOrUpdateUsuario(usuario *Usuario) error
	Update(usuario *Usuario) error
	UpdateUsuarioByEmail(usuario *TipoUsuarioKafka) error
	Delete(usuario *Usuario) error
	GetUsuarioByEMail(email string) (usuario *Usuario, err error)
	GetUsuarioByID(id int) (usuario *Usuario, err error)
	GetPatrimoniosByUsuarioEmail(email string) (patrimonios []uint, err error)
}
