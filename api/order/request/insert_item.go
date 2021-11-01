package request

import (
	"AltaEcom/business/order"
)

type OrderItemRequest struct {
	ID        int `json:"id"`
	ProductID int `json:"product_id"`
	Price     int `json:"price"`
	Qty       int `json:"qty"`
	OrderID   int `json:"order_id"`
}

func (ins OrderItemRequest) ToProductSpec() *order.OrderItemSpec {
	var spec order.OrderItemSpec

	spec.ID = ins.ID
	spec.ProductID = ins.ProductID
	spec.Price = ins.Price
	spec.Qty = ins.Qty
	spec.OrderID = ins.OrderID

	return &spec
}
