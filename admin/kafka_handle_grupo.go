package vertc_go_admin

import (
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

func CreateGrupo(m kafka.Message, usecaseGrupo IUsecaseGrupo,
	usecaseUsuario IUsecaseUsuario) error {

	var grupoJson GrupoJson
	var rupo Grupo

	err := json.Unmarshal(m.Value, &grupoJson)

	if err != nil {
		return err
	}
	rupo.ID = 0
	rupo.IDOPS = grupoJson.IDOPS
	rupo.Nome = grupoJson.Nome
	rupo.Deleted = grupoJson.Deleted
	rupo.Created = grupoJson.Created

	for _, usuarioEmail := range grupoJson.Usuarios {
		usuario, err := usecaseUsuario.GetUsuarioByEMail(usuarioEmail)
		if err != nil {
			return err
		}
		rupo.Usuarios = append(rupo.Usuarios, uint(usuario.ID))
	}

	err = usecaseGrupo.Create(&rupo)

	return err
}

func UpdateGrupo(m kafka.Message, usecaseGrupo IUsecaseGrupo,
	usecaseUsuario IUsecaseUsuario) error {

	log.Println("UPDATE")

	var grupoJson GrupoJson
	var rupo Grupo

	err := json.Unmarshal(m.Value, &grupoJson)

	if err != nil {
		return err
	}

	rupo.IDOPS = grupoJson.IDOPS
	rupo.Nome = grupoJson.Nome

	for _, usuarioEmail := range grupoJson.Usuarios {
		usuario, err := usecaseUsuario.GetUsuarioByEMail(usuarioEmail)
		if err != nil {
			return err
		}
		rupo.Usuarios = append(rupo.Usuarios, uint(usuario.ID))
	}

	_, err = usecaseGrupo.Update(&rupo)

	return err
}

func DeleteGrupo(m kafka.Message, usecaseGrupo IUsecaseGrupo,
	usecaseUsuario IUsecaseUsuario) error {
	var rupo GrupoJson
	log.Println("DELETE")

	err := json.Unmarshal(m.Value, &rupo)

	if err != nil {
		return err
	}

	err = usecaseGrupo.Delete(&rupo)

	return err
}
