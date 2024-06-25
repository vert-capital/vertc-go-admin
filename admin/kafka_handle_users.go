package vertc_go_admin

import (
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

func CreateUser(m kafka.Message, ucUser IUsecaseUsers) error {
	var user UserSSO
	err := json.Unmarshal(m.Value, &user)
	if err != nil {
		return err
	}
	log.Printf("Creating user: %v", user.Email)
	return ucUser.CreateOrUpdateUser(&user)

}
