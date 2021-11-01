package order

import "time"

type OrderItem struct {
	ID          int
	OrderID     int
	ProductID   int
	ProductName string
	Price       int
	Qty         int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Order struct {
	ID        int
	UserID    int
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrderDetail struct {
	ID        int
	UserID    int
	Items     []OrderItem
	UpdatedAt time.Time
}

func NewOrderItem(
	ProductID int,
	OrderID int,
	Price int,
	Qty int,
	CreatedAt time.Time,
	UpdatedAt time.Time) OrderItem {
	return OrderItem{
		ProductID: ProductID,
		OrderID:   OrderID,
		Price:     Price,
		Qty:       Qty,
		CreatedAt: CreatedAt,
		UpdatedAt: UpdatedAt,
	}
}

func ModifyOrderItem(
	ID int,
	ProductID int,
	Qty int,
	Price int,
) OrderItem {
	return OrderItem{
		ID:        ID,
		ProductID: ProductID,
		Price:     Price,
		Qty:       Qty,
		UpdatedAt: time.Now(),
	}
}

func NewOrder(
	UserID int,
	Status bool,
	CreatedAt time.Time) *Order {
	return &Order{
		UserID:    UserID,
		Status:    Status,
		CreatedAt: CreatedAt,
	}
}

func NewOrderDetail(
	Order *Order,
	items *[]OrderItem,
) *OrderDetail {
	var orderDetail OrderDetail

	orderDetail.ID = Order.ID
	orderDetail.UserID = Order.UserID
	orderDetail.UpdatedAt = Order.UpdatedAt

	for _, val := range *items {
		orderDetail.Items = append(orderDetail.Items, val)
	}
	if orderDetail.Items == nil {
		orderDetail.Items = []OrderItem{}
	}

	return &orderDetail
}
