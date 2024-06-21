package vertc_go_admin

import (
	"encoding/json"

	"log"

	"github.com/segmentio/kafka-go"
)

func CreateUsuario(m kafka.Message, usecaseUsuario IUsecaseUsuario) error {

	var usuario Usuario

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

func UpdateTipoUsuario(m kafka.Message, usecaseUsuario IUsecaseUsuario) error {
	var tipoUsuario TipoUsuarioKafka

	err := json.Unmarshal(m.Value, &tipoUsuario)

	if err != nil {
		log.Println("Erro ao deserializar tipo do usuario: ", err)
		return err
	}

	return usecaseUsuario.UpdateUsuarioByEmail(&tipoUsuario)
}
