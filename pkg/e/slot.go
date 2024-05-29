package e

import "net/http"

var (
	ErrSeatReserved = CustomErr{
		HttpStatusCode: http.StatusBadRequest,
		Code:           10000000,
		Msg:            "Seat reserved before",
		Language:       "seat_reserved_before",
	}
	ErrSeatNotInRoomSlot = CustomErr{
		HttpStatusCode: http.StatusBadRequest,
		Code:           10000001,
		Msg:            "Seat not in slot room",
	}
	ErrReserveSeat = CustomErr{
		HttpStatusCode: http.StatusBadRequest,
		Code:           10000002,
		Msg:            "Reserve seat error",
	}
	ErrSeatReserving = CustomErr{
		HttpStatusCode: http.StatusBadRequest,
		Code:           10000003,
		Msg:            "Seat reserving before",
		Language:       "seat_reserved_before",
	}
)
