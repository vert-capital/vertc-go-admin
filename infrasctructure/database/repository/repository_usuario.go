package vertc_go_admin

import (
	entity "github.com/vert-capital/vertc-go-admin/entity"

	"gorm.io/gorm"
)

type RepositoryUsuario struct {
	DB *gorm.DB
}

func NewUsuarioPostgres(DB *gorm.DB) *RepositoryUsuario {
	return &RepositoryUsuario{DB: DB}
}

func (u *RepositoryUsuario) GetByID(id int) (usuario *entity.Usuario, err error) {
	err = u.DB.First(&usuario, id).Error

	if err != nil {
		return nil, err
	}

	return usuario, err
}

func (u *RepositoryUsuario) GetByEMail(email string) (usuario *entity.Usuario, err error) {

	err = u.DB.Where("email = ?", email).First(&usuario).Error

	if err != nil {
		return nil, err
	}

	grupos, err := u.GetUsuarioGrupos(int(usuario.ID))

	if err != nil {
		grupos = make([]entity.Grupo, 0)
	}
	usuario.Grupos = grupos

	return usuario, nil
}

func (u *RepositoryUsuario) CreateUsuario(usuario *entity.Usuario) error {
	return u.DB.Create(&usuario).Error
}

func (u *RepositoryUsuario) Existe(usuario *entity.Usuario) (bool, error) {
	var count int64
	u.DB.Model(&entity.Usuario{}).Where("email = ?", usuario.Email).Count(&count)
	return count > 0, nil
}

func (u *RepositoryUsuario) CreateOrUpdate(usuario *entity.Usuario) error {

	rsp, err := u.Existe(usuario)

	if rsp && err == nil {
		usuario.ID = 0
		return u.DB.Model(&entity.Usuario{}).Where("email = ?", usuario.Email).Updates(usuario).Error
	}

	usuario.ID = 0

	return u.DB.Save(usuario).Error
}

func (u *RepositoryUsuario) UpdateUsuario(usuario *entity.Usuario) error {

	_, err := u.GetByEMail(usuario.Email)

	if err != nil {
		return err
	}

	return u.DB.Save(&usuario).Error
}

func (u *RepositoryUsuario) DeleteUsuario(usuario *entity.Usuario) error {

	_, err := u.GetByEMail(usuario.Email)

	if err != nil {
		return err
	}

	return u.DB.Delete(&usuario).Error
}

func (u *RepositoryUsuario) GetUsuarios(filtros entity.UsuarioFiltros) (usuarios []entity.Usuario, err error) {

	usuarios = make([]entity.Usuario, 0)

	query := u.DB

	if filtros.Busca != "" {
		query = query.Where("nome LIKE ? or email LIKE ?", "%"+filtros.Busca+"%")
	}

	err = query.Find(&usuarios).Error

	return usuarios, err
}

func (u *RepositoryUsuario) GetUsuario(id int) (usuario *entity.Usuario, err error) {

	usuario = &entity.Usuario{}

	err = u.DB.First(&usuario, id).Error

	if err != nil {
		return nil, err
	}
	u.DB.Where("usuario_id = ?", id).Find(&usuario.Grupos)

	return usuario, err
}

func (u *RepositoryUsuario) GetUsuarioGrupos(id int) ([]entity.Grupo, error) {

	var grupos []entity.Grupo

	err := u.DB.Model(&entity.Usuario{ID: id}).Association("Grupos").Find(&grupos)
	if err != nil {
		return nil, err
	}
	return grupos, nil
}

func (u *RepositoryUsuario) GetPatrimoniosByEmail(email string) (patrimonios []uint, err error) {
	err = u.DB.Table("patrimonios").
		Select("patrimonios.id").
		Joins("JOIN responsaveis ON responsaveis.id_emissao = patrimonios.id_emissao").
		Where("responsaveis.email = ?", email).
		Find(&patrimonios).Error

	return patrimonios, err
}

func (u *RepositoryUsuario) UpdateByEmail(usuario *entity.TipoUsuarioKafka) error {
	return u.DB.Model(&entity.Usuario{}).Where("email = ?", usuario.Email).
		Updates(map[string]interface{}{"tipo": usuario.Tipo, "pode_mudar_tipo": usuario.PodeMudarTipo}).Error
}
