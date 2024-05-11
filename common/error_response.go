package common

import (
	"errors"
	"net/http"
)

type ErrorRes struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"key"`
}

func ErrDB(err error) *ErrorRes {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "Something went wrong", err.Error(), "DB_ERROR")
}

func NewErrorResponse(root error, msg, log, key string) *ErrorRes {
	return &ErrorRes{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *ErrorRes {
	return &ErrorRes{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorized(root error, msg, log, key string) *ErrorRes {
	return &ErrorRes{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func ErrorCannotGetEntity(entity string, root error) *ErrorRes {
	return &ErrorRes{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    "Cannot get " + entity,
		Log:        root.Error(),
		Key:        "ErrorCannotGetEntity",
	}
}

func ErrInternal(root error) *ErrorRes {
	return &ErrorRes{
		StatusCode: http.StatusInternalServerError,
		RootErr:    root,
		Message:    "Internal error",
		Log:        root.Error(),
		Key:        "ErrorInternal",
	}
}

func (e *ErrorRes) RootError() error {
	if err, ok := e.RootErr.(*ErrorRes); ok {
		return err.RootError()
	}

	return e.RootErr
}

func (e *ErrorRes) Error() string {
	return e.RootError().Error()
}

var RecordNotFound = errors.New("Record not found")
