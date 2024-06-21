package vertc_go_admin

import entity "github.com/vert-capital/vertc-go-admin/entity"

type IRepositoryGrupo interface {
	GetByID(id int) (grupo *entity.Grupo, err error)
	GetGrupoByIDOPS(idOps int) (*entity.Grupo, error)
	CreateGrupo(grupo *entity.Grupo) error
	CreateGrupoWithUsuarios(grupo *entity.Grupo) error
	UpdateGrupo(grupo *entity.Grupo) (*entity.Grupo, error)
	UpdateGrupoWithUsuarios(grupo *entity.Grupo) error
	DeleteGrupo(grupo *entity.GrupoJson) error
	GetGrupos() (grupos []entity.Grupo, err error)
}

type IUsecaseGrupo interface {
	GetGrupos() (grupos []entity.Grupo, err error)
	Create(grupo *entity.Grupo) error
	Update(grupo *entity.Grupo) (*entity.Grupo, error)
	Delete(grupo *entity.GrupoJson) error
	GetByID(id int) (grupo *entity.Grupo, err error)
}
