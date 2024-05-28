package e

import "net/http"

var (
	ErrSeatReserved = CustomErr{
		HttpStatusCode: http.StatusBadRequest,
		Code:           10000000,
		Msg:            "Seat reserved before",
		Language:       "seat_reserved_before",
	}
)
