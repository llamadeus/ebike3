package utils

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

type (
	ValidationError struct {
		entries []ValidationErrorEntry
	}

	ValidationErrorEntry struct {
		Field string `json:"field"`
		Rule  string `json:"rule"`
	}
)

func (v ValidationError) Error() string {
	return "Validation error"
}

func (v ValidationError) Entries() []ValidationErrorEntry {
	return v.entries
}

func ValidateStruct[T any](value *T) error {
	var validate = validator.New(validator.WithRequiredStructEnabled())

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	err := validate.Struct(value)
	if err != nil {
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			return nil
		}

		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			return formatValidationErrors(validationErrors)
		}

		return err
	}

	return nil
}

func formatValidationErrors(errors validator.ValidationErrors) *ValidationError {
	var result ValidationError

	for _, err := range errors {
		result.entries = append(result.entries, ValidationErrorEntry{
			Field: getFieldPath(err.Namespace()),
			Rule:  err.Tag(),
		})
	}

	return &result
}

func getFieldPath(namespace string) string {
	openBracket := false

	for i := 0; i < len(namespace); i++ {
		switch namespace[i] {
		case '[':
			openBracket = true
		case ']':
			openBracket = false
		case '.':
			if !openBracket {
				return namespace[i+1:]
			}
		}
	}

	return namespace
}
