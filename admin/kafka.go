package vertc_go_admin

import (
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

func StartKafka(db *gorm.DB) {

	// repositorios
	repositoryUsuario := NewUsuarioPostgres(db)
	repositoryGrupo := NewGrupoPostgres(db)

	// usecases
	usecaseUsuario := NewService(repositoryUsuario)
	usecaseGrupo := NewUseCaseGrupo(repositoryGrupo)

	var topicParams []KafkaReadTopicsParams

	topicParams = append(topicParams,
		KafkaReadTopicsParams{
			Topic: "vertc-user",
			Handler: func(m kafka.Message) error {

				log.Println("TOPIC: " + m.Topic)
				log.Println("Message received: ", string(m.Value))

				err := CreateUsuario(m, usecaseUsuario)
				if err != nil {
					log.Printf("Erro ao criar usuário: %v", err)
				}
				return nil
			},
		}, KafkaReadTopicsParams{
			Topic: "vertc-user-group",
			Handler: func(m kafka.Message) error {

				log.Println("TOPIC: " + m.Topic)
				log.Println("Message received: ", string(m.Value))

				var rupo GrupoJson

				err := json.Unmarshal(m.Value, &rupo)
				if err != nil {
					log.Println("Erro ao deserializar o grupo: ", err)
					return nil
				}

				if rupo.Created && !rupo.Deleted {
					err := CreateGrupo(m, usecaseGrupo, usecaseUsuario)
					if err != nil {
						log.Printf("Erro ao criar grupo: %v", err)
					}
				}

				if !rupo.Created && !rupo.Deleted {
					err := UpdateGrupo(m, usecaseGrupo, usecaseUsuario)
					if err != nil {
						log.Printf("Erro ao atualizar grupo: %v", err)
					}
				}

				if rupo.Deleted {
					err := DeleteGrupo(m, usecaseGrupo, usecaseUsuario)
					if err != nil {
						log.Printf("Erro ao deletar grupo: %v", err)
					}
					return nil
				}

				return nil
			},
		},
	)

	startKafkaConnection(topicParams)
	readTopics()
}
