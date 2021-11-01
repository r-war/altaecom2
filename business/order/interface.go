package order

import "time"

type Service interface {
	GetOrderByUserID(userID int) (*Order, error)

	GetOrderItemByUserID(id int) (*OrderDetail, error)

	NewOrderByUserID(userID int) (*Order, error)

	ChangeOrderStatus(id int, ischeckout bool) error

	AddItemToOrder(orderID int, product *OrderItemSpec) error

	UpdateItemInOrder(orderID int, product *OrderItemSpec) error

	RemoveItemInOrder(OrderID int, productID int) error
}

type Repository interface {
	GetOrderByUserID(userID int) (*Order, error)

	NewOrderByUserID(userID int, createdAt time.Time) (*Order, error)

	ChangeOrderStatus(id int, ischekout bool) error
}

type RepositoryOrderItem interface {
	GetOrderItemByUserID(orderID int) (*[]OrderItem, error)

	AddItemToOrder(orderID int, product OrderItem) error

	UpdateItemInOrder(orderID int, product OrderItem) error

	RemoveItemInOrder(orderID int, productID int) error
}
