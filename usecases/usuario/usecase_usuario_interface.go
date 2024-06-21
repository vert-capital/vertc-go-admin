package usecase_usuario

import "github.com/vert-capital/vertc-go-admin/entity"

//go:generate mockgen -destination=../../mocks/mock_usecase_repository_usuario.go -package=mocks app/usecase/usuario IRepositoryUsuario
type IRepositoryUsuario interface {
	GetByID(id int) (usuario *entity.Usuario, err error)
	GetByEMail(email string) (usuario *entity.Usuario, err error)
	CreateUsuario(usuario *entity.Usuario) error
	UpdateUsuario(usuario *entity.Usuario) error
	UpdateByEmail(usuario *entity.TipoUsuarioKafka) error
	DeleteUsuario(usuario *entity.Usuario) error
	GetUsuario(id int) (usuario *entity.Usuario, err error)
	GetUsuarioGrupos(id int) (grupos []entity.Grupo, err error)
	GetPatrimoniosByEmail(email string) (patrimonios []uint, err error)
	CreateOrUpdate(usuario *entity.Usuario) error
	Existe(usuario *entity.Usuario) (bool, error)
}

//go:generate mockgen -destination=../../mocks/mock_usecase_usuario.go -package=mocks app/usecase/usuario IUsecaseUsuario
type IUsecaseUsuario interface {
	GetUsuarioByToken(token string) (*entity.Usuario, error)
	Create(usuario *entity.Usuario) error
	CreateOrUpdateUsuario(usuario *entity.Usuario) error
	Update(usuario *entity.Usuario) error
	UpdateUsuarioByEmail(usuario *entity.TipoUsuarioKafka) error
	Delete(usuario *entity.Usuario) error
	GetUsuarioByEMail(email string) (usuario *entity.Usuario, err error)
	GetUsuarioByID(id int) (usuario *entity.Usuario, err error)
	GetPatrimoniosByUsuarioEmail(email string) (patrimonios []uint, err error)
}
