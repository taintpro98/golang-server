package constants

type SeatStatus string

const (
	ReservingSeat SeatStatus = "reserving"
	ReservedSeat  SeatStatus = "reserved"
)

type SeatType string

const (
	NormalSeat SeatType = "normal"
	VIPSeat    SeatType = "vip"
)
