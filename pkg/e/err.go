package e

import (
	"fmt"
	"net/http"
)

type CustomErr struct {
	HttpStatusCode int    `json:"-"`
	Code           int    `json:"code"`
	Msg            string `json:"message"`
	Language       string `json:"language"`
}

func (c CustomErr) Error() string {
	return c.Msg
}

var (
	ErrNilResponse = CustomErr{
		HttpStatusCode: http.StatusServiceUnavailable,
		Code:           http.StatusServiceUnavailable,
		Msg:            "External service is unavailable",
	}
	ErrUnauthorized = CustomErr{
		HttpStatusCode: http.StatusUnauthorized,
		Code:           http.StatusUnauthorized,
		Msg:            "Unauthorized",
		Language:       "unauthorized",
	}
	ErrTimeout = CustomErr{
		HttpStatusCode: http.StatusRequestTimeout,
		Code:           http.StatusRequestTimeout,
		Msg:            "Request timeout",
	}
)

func ErrDataNotFound(tableName string) CustomErr {
	return CustomErr{
		HttpStatusCode: http.StatusBadRequest,
		Code:           http.StatusNotFound,
		Msg:            fmt.Sprintf("%s not found", tableName),
	}
}
