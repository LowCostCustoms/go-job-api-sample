package errors

type GenericError[T any] struct {
	message string
}

func (e *GenericError[T]) Error() string {
	return e.message
}

type NotFoundTag struct{}
type BadRequestTag struct{}

func NewNotFoundError(message string) error {
	return &GenericError[NotFoundTag]{
		message: message,
	}
}

func NewBadRequestError(message string) error {
	return &GenericError[BadRequestTag]{
		message: message,
	}
}
