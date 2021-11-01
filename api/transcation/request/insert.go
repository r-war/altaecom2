package request

import (
	"AltaEcom/business/transcation"
	"time"
)

type TranscationRequest struct {
	OrderID     int       `json:"order_id"`
	Status      string    `json:"status"`
	Total       int       `json:"total"`
	TypePayment string    `json:"type_payment"`
	DatePayment time.Time `json:"date_payment"`
	Comment     string    `json:"comment"`
}

func (req TranscationRequest) ToTranscationSpec() *transcation.TranscationSpec {
	var spec transcation.TranscationSpec

	spec.OrderID = req.OrderID
	spec.Status = req.Status
	spec.Total = req.Total
	spec.TypePayment = req.TypePayment
	spec.DatePayment = req.DatePayment
	spec.Comment = req.Comment

	return &spec
}
