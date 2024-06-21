package vertc_go_admin

import (
	"log"
	"time"

	"github.com/go-playground/validator/v10"
)

type IError struct {
	Field string
	Tag   string
	Value string
}

var validate *validator.Validate = validator.New()

func GetStructError(err error) []IError {
	var errors []IError

	if err == nil {
		return errors
	}

	for _, err := range err.(validator.ValidationErrors) {
		errors = append(errors, IError{
			Field: err.Field(),
			Tag:   err.Tag(),
			Value: err.Value().(string),
		})
	}

	return errors
}

func FormatDate(date string) (formattedDate time.Time, err error) {
	loc, _ := time.LoadLocation("America/Sao_Paulo")

	var dateFormats = []string{
		time.RFC3339,
		"2006-01-02T16:44:36.611268",
		"2006-01-02 00:00:00 -0300",
		"2006-01-02T15:04:05",
		"2006-01-02",
	}

	var parsedTime time.Time
	var errDate error

	for _, format := range dateFormats {
		parsedTime, errDate = time.ParseInLocation(format, date, loc)
		if errDate == nil {
			break
		}
	}

	if errDate != nil {
		log.Println("Erro ao converter data: ", date)
		return parsedTime, errDate
	}

	return parsedTime, nil
}
