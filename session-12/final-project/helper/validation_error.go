package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func GetValidationErrorMsg(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fieldError.Field())
	case "email":
		return fmt.Sprintf("Invalid %s address", fieldError.Field())
	case "min":
		return fmt.Sprintf("Your %s must be have at least %s characters long", fieldError.Field(), fieldError.Param())
	case "gt":
		return fmt.Sprintf("Should be greater than %s", fieldError.Param())
	}
	return "Unknown error"
}
