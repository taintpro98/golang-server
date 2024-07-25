package constants

const MBatchSize = 100

const (
	TraceID    string = "request_id"
	XUserID    string = "x-user-id"
	XRequestID string = "x-request-id"
)

const (
	SlotSeatsMapKey = "slot_seats_map_key" // reserving seats

	ListRoomsKey     = "list_rooms_key"
	FindByIDSlotKey  = "find_by_id_slot_key"
	ListCacheSlotKey = "list_cache_slot_key"

	PostsChannel    = "posts_channel"
	MessagesChannel = "messages_channel"
)

type GolangServerConstant string

const (
	UsersNum GolangServerConstant = "users_num"
)
