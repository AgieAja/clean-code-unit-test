package validation

import (
	"clean-code-unit-test/model/dto/json"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"github.com/stoewer/go-strcase"
)

func GetvalidationError(err error) []json.ValidationField {
	var validationFields []json.ValidationField
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, validationError := range ve {
			log.Debug().Msg(fmt.Sprintf("validationError : %v", validationError))
			myField := convertFieldRequired(validationError.Namespace())
			validationFields = append(validationFields, json.ValidationField{
				FieldName: myField,
				Message:   formatMessage(validationError),
			})
		}
	}
	return validationFields
}

func formatMessage(err validator.FieldError) string {
	var message string

	switch err.Tag() {
	case "required":
		message = "required"
	case "number":
		message = "must be number"
	case "email":
		message = "invalid format email"
	case "DateOnly":
		message = "invalid format date"
	case "min":
		message = "minimum value is not exceed"
	case "max":
		message = "max value is exceed"
	}

	return message
}

func convertFieldRequired(myValue string) string {
	fieldSegmen := strings.Split(myValue, ".")

	myField := ""
	length := len(fieldSegmen)
	i := 1
	for _, val := range fieldSegmen {
		if i == 1 {
			i++
			continue
		}

		if i == length {
			myField += strcase.SnakeCase(val)
			break
		}

		myField += strcase.LowerCamelCase(val) + `/`
		i++
	}

	return myField
}
