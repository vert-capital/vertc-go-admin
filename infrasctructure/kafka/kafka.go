package kafka

import (
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/vert-capital/vertc-go-admin/entity"
	"github.com/vert-capital/vertc-go-admin/infrasctructure/database/repository"
	kafka_handlers "github.com/vert-capital/vertc-go-admin/infrasctructure/kafka/handlers"
	usecase_grupo "github.com/vert-capital/vertc-go-admin/usecases/grupo"
	usecase_usuario "github.com/vert-capital/vertc-go-admin/usecases/usuario"
	"gorm.io/gorm"
)

func StartKafka(db *gorm.DB) {

	// repositorios
	repositoryUsuario := repository.NewUsuarioPostgres(db)
	repositoryGrupo := repository.NewGrupoPostgres(db)

	// usecases
	usecaseUsuario := usecase_usuario.NewService(repositoryUsuario)
	usecaseGrupo := usecase_grupo.NewUseCaseGrupo(repositoryGrupo)

	var topicParams []KafkaReadTopicsParams

	topicParams = append(topicParams,
		KafkaReadTopicsParams{
			Topic: "vertc-user",
			Handler: func(m kafka.Message) error {

				log.Println("TOPIC: " + m.Topic)
				log.Println("Message received: ", string(m.Value))

				err := kafka_handlers.CreateUsuario(m, usecaseUsuario)
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

				var entityGrupo entity.GrupoJson

				err := json.Unmarshal(m.Value, &entityGrupo)
				if err != nil {
					log.Println("Erro ao deserializar o grupo: ", err)
					return nil
				}

				if entityGrupo.Created && !entityGrupo.Deleted {
					err := kafka_handlers.CreateGrupo(m, usecaseGrupo, usecaseUsuario)
					if err != nil {
						log.Printf("Erro ao criar grupo: %v", err)
					}
				}

				if !entityGrupo.Created && !entityGrupo.Deleted {
					err := kafka_handlers.UpdateGrupo(m, usecaseGrupo, usecaseUsuario)
					if err != nil {
						log.Printf("Erro ao atualizar grupo: %v", err)
					}
				}

				if entityGrupo.Deleted {
					err := kafka_handlers.DeleteGrupo(m, usecaseGrupo, usecaseUsuario)
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
