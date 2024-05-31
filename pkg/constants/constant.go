package constants

const (
	TraceID      = "request_id"
	KeyRequestID = "X-REQUEST-ID"
	XUserID      = "x-user-id"
)

const (
	SlotSeatsMapKey = "slot_seats_map_key" // reserving seats

	ListRoomsKey     = "list_rooms_key"
	FindByIDSlotKey  = "find_by_id_slot_key"
	ListCacheSlotKey = "list_cache_slot_key"
)

type GolangServerConstant string

const (
	UsersNum GolangServerConstant = "users_num"
)
