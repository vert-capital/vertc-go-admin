package vertc_go_admin

import (
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
	entity "github.com/vert-capital/vertc-go-admin/entity"
	usecase_grupo "github.com/vert-capital/vertc-go-admin/usecases/grupo"
	usecase_usuario "github.com/vert-capital/vertc-go-admin/usecases/usuario"
)

func CreateGrupo(m kafka.Message, usecaseGrupo usecase_grupo.IUsecaseGrupo,
	usecaseUsuario usecase_usuario.IUsecaseUsuario) error {

	var grupoJson entity.GrupoJson
	var entityGrupo entity.Grupo

	err := json.Unmarshal(m.Value, &grupoJson)

	if err != nil {
		return err
	}
	entityGrupo.ID = 0
	entityGrupo.IDOPS = grupoJson.IDOPS
	entityGrupo.Nome = grupoJson.Nome
	entityGrupo.Deleted = grupoJson.Deleted
	entityGrupo.Created = grupoJson.Created

	for _, usuarioEmail := range grupoJson.Usuarios {
		usuario, err := usecaseUsuario.GetUsuarioByEMail(usuarioEmail)
		if err != nil {
			return err
		}
		entityGrupo.Usuarios = append(entityGrupo.Usuarios, uint(usuario.ID))
	}

	err = usecaseGrupo.Create(&entityGrupo)

	return err
}

func UpdateGrupo(m kafka.Message, usecaseGrupo usecase_grupo.IUsecaseGrupo,
	usecaseUsuario usecase_usuario.IUsecaseUsuario) error {

	log.Println("UPDATE")

	var grupoJson entity.GrupoJson
	var entityGrupo entity.Grupo

	err := json.Unmarshal(m.Value, &grupoJson)

	if err != nil {
		return err
	}

	entityGrupo.IDOPS = grupoJson.IDOPS
	entityGrupo.Nome = grupoJson.Nome

	for _, usuarioEmail := range grupoJson.Usuarios {
		usuario, err := usecaseUsuario.GetUsuarioByEMail(usuarioEmail)
		if err != nil {
			return err
		}
		entityGrupo.Usuarios = append(entityGrupo.Usuarios, uint(usuario.ID))
	}

	_, err = usecaseGrupo.Update(&entityGrupo)

	return err
}

func DeleteGrupo(m kafka.Message, usecaseGrupo usecase_grupo.IUsecaseGrupo,
	usecaseUsuario usecase_usuario.IUsecaseUsuario) error {
	var entityGrupo entity.GrupoJson
	log.Println("DELETE")

	err := json.Unmarshal(m.Value, &entityGrupo)

	if err != nil {
		return err
	}

	err = usecaseGrupo.Delete(&entityGrupo)

	return err
}
