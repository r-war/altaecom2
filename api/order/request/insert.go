package request

import (
	"AltaEcom/business/order"
)

type OrderRequest struct {
	UserID int `json:"user_id"`
}

func (ins OrderRequest) ToOrderSpec() *order.OrderSpec {
	var spec order.OrderSpec

	spec.UserID = ins.UserID
	spec.Status = false

	return &spec
}
