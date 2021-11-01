package transcation

import "time"

type Status string

const (
	Pending   Status = "pending"
	Confirmed Status = "confirmed"
	Send      Status = "send"
	Delivered Status = "delivered"
	Cancel    Status = "Cancel"
)

type TypePayment string

const (
	Transfer TypePayment = "transfer"
	Credit   TypePayment = "credit"
	COD      TypePayment = "cod"
)

type Transcation struct {
	ID          int
	OrderID     int
	Status      Status
	Total       int
	TypePayment TypePayment
	DatePayment time.Time
	Comment     string
}

func NewTranscation(
	OrderID int,
	StatusTrans string,
	total int,
	TypePay string,
	DatePayment time.Time,
	Comment string,
) Transcation {
	return Transcation{
		OrderID:     OrderID,
		Status:      Status(StatusTrans),
		Total:       total,
		TypePayment: TypePayment(TypePay),
		DatePayment: DatePayment,
		Comment:     Comment,
	}
}

func ModifyTranscation(
	StatusTrans string,
	TypePay string,
	total int,
) Transcation {
	return Transcation{
		Status:      Status(StatusTrans),
		TypePayment: TypePayment(TypePay),
		Total:       total,
	}
}
