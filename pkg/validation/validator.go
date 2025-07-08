package validation

import (
	"fmt"
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
	once     sync.Once
)

// GetValidator returns the global validator instance
func GetValidator() *validator.Validate {
	once.Do(func() {
		validate = validator.New()
	})

	return validate
}

// ValidateStruct validates a struct and returns user-friendly error messages
func ValidateStruct(s interface{}) error {
	return GetValidator().Struct(s)
}

// formatValidationError converts a validator.FieldError to a user-friendly message
func formatValidationError(err validator.FieldError) string {
	field := err.Field()
	tag := err.Tag()
	param := err.Param()

	switch tag {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters long", field, param)
	case "max":
		return fmt.Sprintf("%s must be at most %s characters long", field, param)
	case "email":
		return fmt.Sprintf("%s must be a valid email address", field)
	case "url":
		return fmt.Sprintf("%s must be a valid URL", field)
	case "numeric":
		return fmt.Sprintf("%s must be numeric", field)
	case "alpha":
		return fmt.Sprintf("%s must contain only alphabetic characters", field)
	case "alphanum":
		return fmt.Sprintf("%s must contain only alphanumeric characters", field)
	default:
		return fmt.Sprintf("%s failed validation for tag '%s'", field, tag)
	}
}

// ValidationErrorResponse represents a structured validation error response
type ValidationErrorResponse struct {
	Error   string   `json:"error"`
	Details []string `json:"details"`
}

// GetValidationErrors extracts validation errors and returns them as a slice of strings
func GetValidationErrors(err error) []string {
	var validationErrors []string

	if validatorErrors, ok := err.(validator.ValidationErrors); ok {
		for _, err := range validatorErrors {
			validationErrors = append(validationErrors, formatValidationError(err))
		}
	} else {
		validationErrors = append(validationErrors, err.Error())
	}

	return validationErrors
}
