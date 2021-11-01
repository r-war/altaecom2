package order

import (
	"AltaEcom/business"
	"AltaEcom/business/order"
	"AltaEcom/business/user"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type Order struct {
	ID        int       `gorm:"id;primaryKey;autoIncrement"`
	UserID    int       `gorm:"user_id"`
	Status    bool      `gorm:"status"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
	user      user.User `gorm:"foreignkey:UserID"`
}

func (ord *Order) newOrderTable(userID int, createdAt time.Time) {
	ord.UserID = userID
	ord.Status = false
	ord.CreatedAt = createdAt

}

func (col *Order) ToOrder() *order.Order {
	return &order.Order{
		ID:        col.ID,
		UserID:    col.UserID,
		Status:    col.Status,
		CreatedAt: col.CreatedAt,
		UpdatedAt: col.UpdatedAt,
	}
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (repo *Repository) GetOrderByUserID(userID int) (*order.Order, error) {
	var cart Order

	err := repo.DB.First(&cart, "status = false and user_id = ? ", userID).Error

	if err != nil {
		return nil, err
	}

	return cart.ToOrder(), nil
}

func (repo *Repository) NewOrderByUserID(userID int, createdAt time.Time) (*order.Order, error) {
	var cart Order

	err := repo.DB.First(&cart, "status = false and user_id = ? ", userID).Error

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, business.ErrHasBeenModified
	}

	cart.newOrderTable(userID, createdAt)
	if err := repo.DB.Create(&cart).Error; err != nil {
		return nil, err
	}

	return cart.ToOrder(), nil
}

func (repo *Repository) ChangeOrderStatus(id int, ischeckout bool) error {
	var cart Order
	err := repo.DB.First(&cart, "user_id = ? ", id).Error

	if err != nil {
		return err
	}
	fmt.Println(&cart)
	err = repo.DB.Model(&cart).Updates(Order{
		Status: ischeckout,
	}).Error

	if err != nil {
		return err
	}

	return nil
}
