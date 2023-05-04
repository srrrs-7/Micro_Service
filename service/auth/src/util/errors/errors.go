package errors

import "errors"

var (
	ErrBadRequest          = errors.New("bad request")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrForbidden           = errors.New("forbidden")
	ErrNotFound            = errors.New("not found")
	ErrInternalServerError = errors.New("internal server error")
	// その他のエラーレスポンスを追加する場合はここに定義する
)
