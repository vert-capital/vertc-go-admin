package vertc_go_admin

import (
	"log"

	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

func StartKafka(db *gorm.DB) {

	// repositorios
	repositoryUsuario := NewRepositoryUsers(db)

	// usecases
	usecaseUsuario := NewUserService(repositoryUsuario)

	var topicParams []KafkaReadTopicsParams

	topicParams = append(topicParams,
		KafkaReadTopicsParams{
			Topic: "vertc-user",
			Handler: func(m kafka.Message) error {

				log.Println("TOPIC: " + m.Topic)
				log.Println("Message received: ", string(m.Value))

				err := CreateUser(m, usecaseUsuario)
				if err != nil {
					log.Printf("Erro ao criar usuário: %v", err)
				}
				return nil
			},
		},
	)

	startKafkaConnection(topicParams)
	readTopics()
}
