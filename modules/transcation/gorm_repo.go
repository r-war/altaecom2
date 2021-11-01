package transcation

import (
	"AltaEcom/business/transcation"
	"time"

	"gorm.io/gorm"
)

type Transcation struct {
	ID          int                     `gorm:"id;primaryKey;autoIncrement"`
	OrderID     int                     `gorm:"order_id"`
	Status      transcation.Status      `gorm:"status"`
	Total       int                     `gorm:"total"`
	TypePayment transcation.TypePayment `gorm:"type_payment"`
	DatePayment time.Time               `gorm:"date_payment"`
	Comment     string                  `gorm:"comment"`
}
type Repository struct {
	DB *gorm.DB
}

func NewTranscation(transcation transcation.Transcation) *Transcation {
	return &Transcation{
		transcation.ID,
		transcation.OrderID,
		transcation.Status,
		transcation.Total,
		transcation.TypePayment,
		transcation.DatePayment,
		transcation.Comment,
	}
}

func (col *Transcation) ToTranscation() transcation.Transcation {
	return transcation.Transcation{
		ID:          col.ID,
		OrderID:     col.OrderID,
		Status:      col.Status,
		Total:       col.Total,
		TypePayment: col.TypePayment,
		DatePayment: col.DatePayment,
		Comment:     col.Comment,
	}
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetAllTranscation() ([]transcation.Transcation, error) {
	var transcations []Transcation

	err := r.DB.Find(&transcations).Error

	if err != nil {
		return nil, err
	}
	var result []transcation.Transcation
	for _, val := range transcations {
		result = append(result, val.ToTranscation())
	}
	return result, nil
}

func (r *Repository) GetTranscationById(id int) (*transcation.Transcation, error) {
	var trans Transcation

	err := r.DB.First(&trans, "id = ? ", id).Error
	if err != nil {
		return nil, err
	}
	result := trans.ToTranscation()

	return &result, nil
}

func (r *Repository) CreateTranscation(transcation transcation.Transcation) error {
	trans := NewTranscation(transcation)

	if err := r.DB.Create(trans).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateTranscation(ID int, transcation transcation.Transcation) error {
	var data Transcation
	err := r.DB.Find(&data, "id = ? ", ID).Error

	if err != nil {
		return err
	}

	err = r.DB.Model(&data).Updates(Transcation{
		Status:      transcation.Status,
		TypePayment: transcation.TypePayment,
		Total:       transcation.Total,
		DatePayment: time.Now(),
	}).Error

	if err != nil {
		return err
	}

	return nil
}
