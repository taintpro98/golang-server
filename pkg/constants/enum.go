package constants

type SeatStatus string

const (
	EmptySeat     SeatStatus = "empty"
	ReservingSeat SeatStatus = "reserving" // trang thai nay chi co trong redis, 2 trang thai con lai luu trong db
	ReservedSeat  SeatStatus = "reserved"
)

type SeatType string

const (
	NormalSeat SeatType = "normal"
	VIPSeat    SeatType = "vip"
)
