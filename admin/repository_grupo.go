package vertc_go_admin

import (
	"log"

	"gorm.io/gorm"
)

type RepositoryGrupo struct {
	DB *gorm.DB
}

func NewGrupoPostgres(DB *gorm.DB) *RepositoryGrupo {
	return &RepositoryGrupo{DB: DB}
}

func (u *RepositoryGrupo) GetByID(id int) (grupo *Grupo, err error) {
	u.DB.First(&grupo, id)

	return grupo, err
}

func (u *RepositoryGrupo) CreateGrupo(grupo *Grupo) error {
	return u.DB.Create(&grupo).Error
}

func (u *RepositoryGrupo) CreateGrupoWithUsuarios(grupo *Grupo) error {
	err := u.CreateGrupo(grupo)
	if err != nil {
		return err
	}

	// Save the usuario-grupo relationships
	for _, usuarioID := range grupo.Usuarios {
		usuarioGrupo := &UsuarioGrupo{
			UsuarioID: usuarioID,
			GrupoID:   uint(grupo.ID),
		}
		err = u.DB.Create(&usuarioGrupo).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *RepositoryGrupo) GetGrupoByIDOPS(idOps int) (*Grupo, error) {
	existingGrupo := &Grupo{}

	if err := u.DB.Table("grupos").Where("id_ops = ?", idOps).First(existingGrupo).Error; err != nil {
		return nil, err
	}
	return existingGrupo, nil
}

func (u *RepositoryGrupo) UpdateGrupo(grupo *Grupo) (*Grupo, error) {
	existingGrupo, err := u.GetGrupoByIDOPS(grupo.IDOPS)
	if err != nil {
		return nil, err
	}

	existingGrupo.Nome = grupo.Nome
	existingGrupo.IDOPS = grupo.IDOPS

	err = u.DB.Where("grupo_id = ?", existingGrupo.ID).Delete(&UsuarioGrupo{}).Error
	if err != nil {
		log.Println("Error update usuario_grupo: ", err)
	}

	if err := u.DB.Save(existingGrupo).Error; err != nil {
		return nil, err
	}

	return existingGrupo, nil
}

func (u *RepositoryGrupo) UpdateGrupoWithUsuarios(grupo *Grupo) error {
	existingGrupo, err := u.UpdateGrupo(grupo)

	if err != nil {
		return err
	}

	err = u.DB.Where("grupo_id = ?", existingGrupo.ID).Delete(&UsuarioGrupo{}).Error
	if err != nil {
		log.Println("Error update usuarios_grupo: ", err)
	}

	for _, usuarioID := range grupo.Usuarios {
		usuarioGrupo := &UsuarioGrupo{
			UsuarioID: usuarioID,
			GrupoID:   uint(existingGrupo.ID),
		}
		err = u.DB.Create(&usuarioGrupo).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *RepositoryGrupo) DeleteGrupo(grupo *GrupoJson) error {

	existingGrupo, err := u.GetGrupoByIDOPS(grupo.IDOPS)
	if err != nil {
		return err
	}

	err = u.DB.Where("grupo_id = ?", existingGrupo.ID).Delete(&UsuarioGrupo{}).Error
	if err != nil {
		log.Println("Error deleting usuario_grupo: ", err)
	}

	return u.DB.Delete(existingGrupo).Error
}

func (u *RepositoryGrupo) GetGrupos() (grupos []Grupo, err error) {
	return grupos, u.DB.Find(&grupos).Error
}
