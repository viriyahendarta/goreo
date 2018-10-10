package error

import "net/http"

type ErrorTypeCode struct {
	HttpCode int
	Message  string
}

const (
	ErrorDefault int = 0

	ErrorInvalidAttribute      int = 100
	ErrorInvalidAuthentication int = 101

	ErrorQuery int = 200
)

func newErrorTypeCode(code int, message string) *ErrorTypeCode {
	return &ErrorTypeCode{
		HttpCode: code,
		Message:  message,
	}
}

var mappingTypeCodeValue = map[int]*ErrorTypeCode{
	ErrorDefault:               newErrorTypeCode(http.StatusInternalServerError, "Default error message"),
	ErrorInvalidAttribute:      newErrorTypeCode(http.StatusBadRequest, "Invalid attribute"),
	ErrorInvalidAuthentication: newErrorTypeCode(http.StatusUnauthorized, "Invalid authentication"),

	ErrorQuery: newErrorTypeCode(http.StatusInternalServerError, "Failed query"),
}

func GetCodeValue(code int) *ErrorTypeCode {
	if v, ok := mappingTypeCodeValue[code]; ok {
		return v
	}
	return mappingTypeCodeValue[ErrorDefault]
}
