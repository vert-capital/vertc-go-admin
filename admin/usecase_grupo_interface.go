package vertc_go_admin

type IRepositoryGrupo interface {
	GetByID(id int) (grupo *Grupo, err error)
	GetGrupoByIDOPS(idOps int) (*Grupo, error)
	CreateGrupo(grupo *Grupo) error
	CreateGrupoWithUsuarios(grupo *Grupo) error
	UpdateGrupo(grupo *Grupo) (*Grupo, error)
	UpdateGrupoWithUsuarios(grupo *Grupo) error
	DeleteGrupo(grupo *GrupoJson) error
	GetGrupos() (grupos []Grupo, err error)
}

type IUsecaseGrupo interface {
	GetGrupos() (grupos []Grupo, err error)
	Create(grupo *Grupo) error
	Update(grupo *Grupo) (*Grupo, error)
	Delete(grupo *GrupoJson) error
	GetByID(id int) (grupo *Grupo, err error)
}
