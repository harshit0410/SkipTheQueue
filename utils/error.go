package utils

import "fmt"

type Error interface {
	Error() string
	ErrorCode() string
	ErrorType() string
}

type GenericError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Type    string `json:"type"`
}

func New(msg string, args ...interface{}) error {
	return fmt.Errorf(msg, args...)
}

func (e *GenericError) Error() string {
	return e.Message
}

func (e *GenericError) ErrorCode() string {
	return e.Code
}

func (e *GenericError) ErrorType() string {
	return e.Type
}
