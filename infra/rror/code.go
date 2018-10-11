package rror

import "net/http"

//Code as error code
type Code int

//Type as error type
type Type string

//Error code initialization
const (
	CodeUnknown Code = 0

	CodeBadRequest          Code = http.StatusBadRequest
	CodeUnauthorized        Code = http.StatusUnauthorized
	CodeForbidden           Code = http.StatusForbidden
	CodeInternalServerError Code = http.StatusInternalServerError
	CodeNotImplemented      Code = http.StatusNotImplemented

	//Custom error code must be 1000 and above
	//below that will be treated as http code

	CodeGeneralQuery Code = 1000
)

var mappingTypeCodeValue = map[Code]Type{
	CodeUnknown: "Unknown error",

	CodeBadRequest:          "Bad request",
	CodeUnauthorized:        "Unauthorized",
	CodeForbidden:           "Forbidden",
	CodeInternalServerError: "Internal Server Error",
	CodeNotImplemented:      "Not Implemented",

	CodeGeneralQuery: "Failed to execute query",
}
