package micro

import "fmt"

type baseError struct {
	StatusCode int
	Message    string
}

var (
	_ error = baseError{}
)

func NewError(code int, message string) error {
	return &baseError{
		StatusCode: code,
		Message:    message,
	}
}

func NewBadRequestError(message string) error {
	return &baseError{
		StatusCode: 400,
		Message:    message,
	}
}

func NewUnauthorizedError(message string) error {
	return &baseError{
		StatusCode: 401,
		Message:    message,
	}
}

func NewForbiddenError(message string) error {
	return &baseError{
		StatusCode: 403,
		Message:    message,
	}
}

func NewNotFoundError(message string) error {
	return &baseError{
		StatusCode: 404,
		Message:    message,
	}
}

func NewInternalServerError(message string) error {
	return &baseError{
		StatusCode: 500,
		Message:    message,
	}
}

func NewNotImplementedError(message string) error {
	return &baseError{
		StatusCode: 501,
		Message:    message,
	}
}

func (e baseError) Error() string {
	return fmt.Sprintf("%d: %s", e.StatusCode, e.Message)
}
