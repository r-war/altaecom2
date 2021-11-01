package response

import (
	"AltaEcom/business/order"
)

type OrderResponse struct {
	ID     int  `json:"id"`
	UserID int  `json:"user_id"`
	Status bool `json:"status"`
}

func NewGetOrderResponse(order *order.Order) *OrderResponse {
	return &OrderResponse{
		ID:     order.ID,
		UserID: order.UserID,
		Status: order.Status,
	}
}
