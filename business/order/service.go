package order

import (
	"AltaEcom/business"
	"AltaEcom/util/validator"
	"fmt"
	"time"
)

type OrderSpec struct {
	UserID int
	Status bool
}

type OrderItemSpec struct {
	ID        int
	ProductID int
	OrderID   int
	Price     int
	Qty       int
}

type service struct {
	repo          Repository
	repoOrderItem RepositoryOrderItem
}

func NewService(repo Repository, repoOrderItem RepositoryOrderItem) Service {
	return &service{repo, repoOrderItem}
}

func (s *service) GetOrderByUserID(userID int) (*Order, error) {
	return s.repo.GetOrderByUserID(userID)
}

func (s *service) NewOrderByUserID(userID int) (*Order, error) {
	return s.repo.NewOrderByUserID(userID, time.Now())
}

func (s *service) GetOrderItemByUserID(id int) (*OrderDetail, error) {
	cart, err := s.repo.GetOrderByUserID(id)

	if err != nil {
		return nil, err
	}

	allItems, err := s.repoOrderItem.GetOrderItemByUserID(id)
	if err != nil {
		return nil, err
	}

	items := NewOrderDetail(cart, allItems)

	return items, nil
}

func (s *service) AddItemToOrder(orderID int, product *OrderItemSpec) error {
	if err := validator.GetValidator().Struct(product); err != nil {
		return business.ErrInvalidSpec
	}

	data := NewOrderItem(
		product.ProductID,
		product.OrderID,
		product.Price,
		product.Qty,
		time.Now(),
		time.Now())
	return s.repoOrderItem.AddItemToOrder(orderID, data)
}

func (s *service) UpdateItemInOrder(orderID int, product *OrderItemSpec) error {
	if err := validator.GetValidator().Struct(product); err != nil {
		return business.ErrInvalidSpec
	}

	fmt.Println(product)
	data := ModifyOrderItem(
		product.ID,
		product.ProductID,
		product.Qty,
		product.Price,
	)

	return s.repoOrderItem.UpdateItemInOrder(orderID, data)
}

func (s *service) RemoveItemInOrder(orderID int, productID int) error {
	return s.repoOrderItem.RemoveItemInOrder(orderID, productID)
}

func (s *service) ChangeOrderStatus(id int, ischeckout bool) error {
	return s.repo.ChangeOrderStatus(id, ischeckout)
}
