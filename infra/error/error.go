package error

import (
	"context"
)

type Error struct {
	ctx     context.Context
	code    int
	message string
	err     error
}

func New(ctx context.Context, code int, message string, err error) error {
	return &Error{
		ctx:     ctx,
		code:    code,
		message: message,
		err:     err,
	}
}

func Cast(err error) *Error {
	if c, ok := err.(*Error); ok {
		return c
	}
	return &Error{
		ctx:     context.Background(),
		code:    0,
		message: "",
		err:     err,
	}
}

func (e *Error) Error() string {
	if e.err != nil {
		return e.err.Error()
	}
	if e.message == "" {
		return e.message
	}
	return "empty error"
}

func (e *Error) GetErrorTypeCode() *ErrorTypeCode {
	return GetCodeValue(e.code)
}

func (e *Error) GetMessage() string {
	return e.message
}
