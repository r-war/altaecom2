package transcation

import (
	"AltaEcom/business"
	"AltaEcom/business/order"
	"AltaEcom/util/validator"
	"time"
)

type TranscationSpec struct {
	Status      string
	TypePayment string
	Total       int
	DatePayment time.Time
	OrderID     int
	Comment     string
}

type service struct {
	repository   Repository
	orderService order.Service
}

func NewService(repo Repository, orderService order.Service) Service {
	return &service{repo, orderService}
}

func (s *service) GetAllTranscation() ([]Transcation, error) {
	transcations, err := s.repository.GetAllTranscation()
	if err != nil {
		return []Transcation{}, err
	}

	return transcations, err
}

func (s *service) GetTranscationById(id int) (*Transcation, error) {
	return s.repository.GetTranscationById(id)
}

func (s *service) CreateTranscation(userID int, trans *TranscationSpec) error {
	var err error
	if err := validator.GetValidator().Struct(trans); err != nil {
		return business.ErrInvalidSpec
	}

	err = s.orderService.ChangeOrderStatus(userID, true)
	if err != nil {
		return business.ErrNotFound
	}

	data := NewTranscation(
		trans.OrderID,
		trans.Status,
		trans.Total,
		trans.TypePayment,
		trans.DatePayment,
		trans.Comment,
	)

	err = s.repository.CreateTranscation(data)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateTranscation(id int, transcation *TranscationSpec) error {
	_, err := s.repository.GetTranscationById(id)

	if err != nil {
		return err
	}
	if err := validator.GetValidator().Struct(transcation); err != nil {
		return business.ErrInvalidSpec
	}

	data := ModifyTranscation(
		transcation.Status,
		transcation.TypePayment,
		transcation.Total,
	)

	return s.repository.UpdateTranscation(id, data)
}
