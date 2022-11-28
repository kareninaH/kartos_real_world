package error

import (
	"fmt"
	"net/http"
)

// 描述: 自定义error
// 作者: hgy
// 创建日期: 2022/11/28

type HttpError struct {
	Code   int                 `json:"-"`
	Errors map[string][]string `json:"errors,omitempty"`
}

func NewHttpError(code int, filed string, message ...string) *HttpError {
	e := HttpError{
		Code: code,
		Errors: map[string][]string{
			filed: message,
		},
	}
	return &e
}

func (h HttpError) Error() string {
	return fmt.Sprintf("HttpError code: %d, body: %s", h.Code, h.Errors)
}

func HttpBadRequest(filed string, message ...string) *HttpError {
	return NewHttpError(http.StatusBadRequest, filed, message...)
}

func HttpUnauthorized(filed string, message ...string) *HttpError {
	return NewHttpError(http.StatusUnauthorized, filed, message...)
}
