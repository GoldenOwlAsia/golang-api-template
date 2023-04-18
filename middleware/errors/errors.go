package errors

import (
	"errors"
	"fmt"
)

type ErrorCode string

const (
	ErrorCodeUnknown        ErrorCode = "UNKNOWN"
	ErrorCodeNotFound       ErrorCode = "NOT_FOUND"
	ErrorCodeBadRequest     ErrorCode = "BAD_REQUEST"
	ErrorCodeUnauthorized   ErrorCode = "UNAUTHORIZED"
	ErrorCodeInvalidRequest ErrorCode = "INVALID_REQUEST"
)

type ErrorType string

const (
	ErrorTypeBadRequest       ErrorType = "BAD_REQUEST"
	ErrorTypeInternal         ErrorType = "INTERNAL"
	ErrorTypeConflict         ErrorType = "CONFLICT"
	ErrorTypeNotFound         ErrorType = "NOT_FOUND"
	ErrorTypeUnauthorized     ErrorType = "UNAUTHORIZED"
	ErrorTypePermissionDenied ErrorType = "PERMISSION_DENIED"
)

type Error struct {
	Code     ErrorCode              `json:"code"`
	Type     ErrorType              `json:"type"`
	Message  string                 `json:"message"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

func New(c ErrorCode, t ErrorType, msg string, md map[string]interface{}) error {
	return Error{
		Code:     c,
		Type:     t,
		Message:  msg,
		Metadata: md,
	}
}

func (e Error) Error() string {
	return fmt.Sprintf("code:%s - type:%s - msg:%s - metadata:%v", e.Code, e.Type, e.Message, e.Metadata)
}

func Convert(err error) (Error, bool) {
	if err == nil {
		panic("err must be non-nil")
	}

	if e := new(Error); errors.As(err, e) {
		return *e, true
	}

	return Error{
		Code:    ErrorCodeUnknown,
		Type:    ErrorTypeInternal,
		Message: fmt.Sprintf("unknown error: %v", err),
	}, false
}
