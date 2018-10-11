package rror

import (
	"fmt"
)

//Error object
type Error struct {
	Code    Code
	Type    Type
	Message string
	Err     error
}

//New creates new Error instance
func New(code Code, message string, err error) error {
	return &Error{
		Code:    code,
		Type:    mappingTypeCodeValue[code],
		Message: message,
		Err:     err,
	}
}

//Cast error interface into Error
func Cast(err error) *Error {
	if c, ok := err.(*Error); ok {
		return c
	}

	return &Error{
		Code:    CodeUnknown,
		Type:    mappingTypeCodeValue[CodeUnknown],
		Message: "",
		Err:     err,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("[!%v] %s: %s (%s)", e.Code, e.Type, e.Message, e.Err.Error())
}
