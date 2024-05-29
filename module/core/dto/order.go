package dto

type CreateOrderRequest struct {
	SlotID  string   `json:"slot_id" binding:"required"`
	SeatIDs []string `json:"seat_ids" binding:"required"`
}

type CreateOrderResponse struct {
	OrderID string `json:"order_id" binding:"required"`
}
