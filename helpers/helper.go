package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

type CustomValidator struct {
	validate *validator.Validate
}

func NewValidator() *CustomValidator {
	validator := validator.New()
	return &CustomValidator{
		validate: validator,
	}
}

func (this *CustomValidator) Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := this.validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = customeTag(err)     // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func customeTag(fl validator.FieldError) string {
	fieldName := fl.Field()
	tagName := fl.Tag()

	// Menyimpan informasi tentang tag dan pesan kesalahan yang sesuai
	var customResponse string
	switch tagName {
	case "required":
		customResponse = fmt.Sprintf("The %s field is required.", fieldName)
	case "max":
		maxLength := fl.Param()
		customResponse = fmt.Sprintf("The %s field must be less than or equal to %s characters.", fieldName, maxLength)
	case "email":
		customResponse = fmt.Sprintf("The %s field must be a valid email address.", fieldName)
	case "min":
		minLength := fl.Param()
		customResponse = fmt.Sprintf("The %s field must be at least %s characters.", fieldName, minLength)
	case "confirmed":
		customResponse = fmt.Sprintf("The %s field must be confirmed.", fieldName)
	case "oneof":
		allowedValues := fl.Param()
		customResponse = fmt.Sprintf("The %s field must be one of %s.", fieldName, allowedValues)
	default:
		customResponse = tagName
	}

	return customResponse
}
