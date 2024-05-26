package dto

type FilterRoom struct {
	ID           int64
	CommonFilter CommonFilter
}

type AdminCreateRoomRequest struct {
	Name string `json:"name"`
}

type AdminCreateRoomResponse struct {
	RoomID int64 `json:"room_id"`
}
