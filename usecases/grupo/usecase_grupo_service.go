package vertc_go_admin

import (
	entity "github.com/vert-capital/vertc-go-admin/entity"
)

type UseCaseGrupo struct {
	repository IRepositoryGrupo
}

func NewUseCaseGrupo(repository IRepositoryGrupo) IUsecaseGrupo {
	return &UseCaseGrupo{repository: repository}
}

func (u *UseCaseGrupo) GetGrupos() (grupos []entity.Grupo, err error) {
	return u.repository.GetGrupos()
}

func (u *UseCaseGrupo) Create(grupo *entity.Grupo) error {
	err := grupo.Validate()

	if err != nil {
		return err
	}
	if len(grupo.Usuarios) == 0 {
		return u.repository.CreateGrupo(grupo)
	}

	return u.repository.CreateGrupoWithUsuarios(grupo)
}

func (u *UseCaseGrupo) Update(grupo *entity.Grupo) (*entity.Grupo, error) {
	err := grupo.Validate()

	if err != nil {
		return nil, err
	}

	if len(grupo.Usuarios) == 0 {
		return u.repository.UpdateGrupo(grupo)
	}

	return nil, u.repository.UpdateGrupoWithUsuarios(grupo)
}

func (u *UseCaseGrupo) Delete(grupo *entity.GrupoJson) error {
	return u.repository.DeleteGrupo(grupo)
}

func (u *UseCaseGrupo) GetByID(id int) (grupo *entity.Grupo, err error) {
	return u.repository.GetByID(id)
}
