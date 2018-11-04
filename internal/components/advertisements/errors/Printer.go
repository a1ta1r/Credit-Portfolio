package errors

import (
	"fmt"
	"gopkg.in/go-playground/validator.v8"
)

func GetErrorMessages(errors validator.ValidationErrors) []string {
	var errorMsg []string
	for _, err := range errors {
		errorMsg = append(errorMsg, getMsg(err))
	}
	return errorMsg
}

func getMsg(error *validator.FieldError) string {
	switch error.Tag {
	case "required":
		return fmt.Sprintf("%s is required", error.Field)
	case "max":
		return fmt.Sprintf("%s cannot be longer than %s", error.Field, error.Param)
	case "min":
		return fmt.Sprintf("%s must be longer than %s", error.Field, error.Param)
	case "email":
		return fmt.Sprintf("Invalid email format")
	case "len":
		return fmt.Sprintf("%s must be %s characters long", error.Field, error.Param)
	}
	return fmt.Sprintf("%s is not valid", error.Field)
}