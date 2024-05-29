package e

import "net/http"

var (
	ErrInvalidOrder = CustomErr{
		HttpStatusCode: http.StatusBadRequest,
		Code:           20000000,
		Msg:            "Invalid order",
	}
	ErrNotReserveSeatBefore = CustomErr{
		HttpStatusCode: http.StatusBadRequest,
		Code:           20000001,
		Msg:            "Not reserve seat before",
	}
)
