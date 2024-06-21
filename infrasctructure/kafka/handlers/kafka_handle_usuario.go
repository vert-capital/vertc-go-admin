package kafka_handlers

import (
	"encoding/json"

	"log"

	"github.com/segmentio/kafka-go"
	"github.com/vert-capital/vertc-go-admin/entity"
	usecase_usuario "github.com/vert-capital/vertc-go-admin/usecases/usuario"
)

func CreateUsuario(m kafka.Message, usecaseUsuario usecase_usuario.IUsecaseUsuario) error {

	var usuario entity.Usuario

	err := json.Unmarshal(m.Value, &usuario)

	if usuario.Imagem == nil {
		var defaultImage string = ""
		usuario.Imagem = &defaultImage
	}

	if err != nil {
		log.Println("Erro ao deserializar o usuario: ", err)
		return err
	}

	log.Printf("Usuario %s recebido via kafka", usuario.Email)
	return usecaseUsuario.CreateOrUpdateUsuario(&usuario)
}

func UpdateTipoUsuario(m kafka.Message, usecaseUsuario usecase_usuario.IUsecaseUsuario) error {
	var tipoUsuario entity.TipoUsuarioKafka

	err := json.Unmarshal(m.Value, &tipoUsuario)

	if err != nil {
		log.Println("Erro ao deserializar tipo do usuario: ", err)
		return err
	}

	return usecaseUsuario.UpdateUsuarioByEmail(&tipoUsuario)
}
