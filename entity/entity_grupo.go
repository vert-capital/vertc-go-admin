package vertc_go_admin

type Grupo struct {
	ID       int    `json:"-" gorm:"primaryKey"`
	IDOPS    int    `json:"id_ops"`
	Nome     string `json:"nome" validate:"required,min=3,max=120"`
	Usuarios []uint `json:"-" gorm:"-"`
	Deleted  bool   `json:"-" gorm:"-"`
	Created  bool   `json:"-" gorm:"-"`
}

type GrupoJson struct {
	IDOPS    int      `json:"id"`
	Nome     string   `json:"name"`
	Usuarios []string `json:"users"`
	Deleted  bool     `json:"deleted"`
	Created  bool     `json:"created"`
}

type UsuarioGrupo struct {
	UsuarioID uint
	GrupoID   uint
}

func (g *Grupo) Validate() error {
	return validate.Struct(g)
}

func (Grupo) TableName() string {
	return "vertadmin_grupos"
}

func (UsuarioGrupo) TableName() string {
	return "vertadmin_usuarios_grupos"
}
