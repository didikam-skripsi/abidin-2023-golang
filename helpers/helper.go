package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func FormatValidationError(err error) []ValidationError {
	var errors []ValidationError

	for _, e := range err.(validator.ValidationErrors) {
		field := e.Field()
		message := fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", field, e.Tag())
		errors = append(errors, ValidationError{
			Field:   field,
			Message: message,
		})
	}

	return errors
}