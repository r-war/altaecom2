package response

import (
	"AltaEcom/business/transcation"
	"time"
)

type TranscationResponse struct {
	ID          int       `json:"id"`
	OrderID     int       `json:"order_id"`
	Status      string    `json:"status"`
	Total       int       `json:"total"`
	TypePayment string    `json:"type_payment"`
	DatePayment time.Time `json:"date_payment"`
	Comment     string    `json:"comment"`
}

func NewTranscationResponse(trans transcation.Transcation) *TranscationResponse {
	var respon TranscationResponse

	respon.ID = trans.ID
	respon.OrderID = trans.OrderID
	respon.Status = string(trans.Status)
	respon.Total = trans.Total
	respon.TypePayment = string(trans.TypePayment)
	respon.DatePayment = trans.DatePayment
	respon.Comment = trans.Comment

	return &respon
}
