package main

import (
	"fmt"
)

// Error codes
const (
	ErrCodeNotExist      = 1
	ErrCodeAlreadyExists = 2
)

// Error The serializable Error structure.
type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

// NewError creates an error instance with the specified code and message.
func NewError(code int, msg string) *Error {
	return &Error{
		Code:    code,
		Message: msg,
	}
}

// Success 成功
type Success struct {
	Result interface{} `json:"result"`
	Code   int
}

// NewSuccess 新しいSuccessを作成
func NewSuccess(event *Event) *Success {
	return &Success{
		Result: event,
		Code:   200,
	}
}

// MajorMessage MajorとMessageを返す
type MajorMessage struct {
	Major   int
	Message string
}
