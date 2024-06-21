package vertc_go_admin

type Usuario struct {
	ID            int     `json:"id"`
	PrimeiroNome  *string `json:"first_name"       validate:"required,min=3,max=120"`
	UltimoNome    *string `json:"last_name"`
	Area          *uint   `json:"area"`
	Email         *string `json:"email"      validate:"required,email"`
	Imagem        *string `json:"image"`
	Tipo          *string `json:"user_type"`
	PodeMudarTipo *bool   `json:"can_change_user_type" gorm:"default:false"`
	Grupos        []Grupo `json:"groups" gorm:"many2many:usuarios_grupos;" on_delete:"cascade"`
	IsAdmin       *bool   `json:"is_superuser" gorm:"default:false"`
}

type UsuarioJson struct {
	ID           int     `json:"id"`
	PrimeiroNome string  `json:"primeiro_nome"`
	UltimoNome   string  `json:"ultimo_nome"`
	Area         uint    `json:"area"`
	Email        string  `json:"email"`
	Imagem       *string `json:"imagem"`
	Grupos       []Grupo `json:"grupos"`
}

type TipoUsuarioKafka struct {
	Email         string `json:"email"`
	Tipo          string `json:"user_type"`
	PodeMudarTipo bool   `json:"can_change_user_type"`
}

type UsuarioFiltros struct {
	IDs   []uint `json:"ids"`
	Busca string `json:"search"`
}

func (Usuario) TableName() string {
	return "vertadmin_usuarios"
}

func (u *Usuario) RespostaJson() UsuarioJson {
	return UsuarioJson{
		ID:           u.ID,
		PrimeiroNome: *u.PrimeiroNome,
		UltimoNome:   *u.UltimoNome,
		Area:         *u.Area,
		Email:        *u.Email,
		Imagem:       u.Imagem,
		Grupos:       u.Grupos,
	}
}

func NewUsuario(usuarioParam Usuario) (*Usuario, error) {

	u := &Usuario{
		Email: usuarioParam.Email,
	}

	return u, nil
}

func (u *Usuario) Validate() error {
	return validate.Struct(u)
}

type AreaUsuario struct {
	ID     uint   `json:"id"`
	Area   string `json:"area"`
	Status string `json:"status"`
}

func (AreaUsuario) TableName() string {
	return "vertadmin_areas"
}
