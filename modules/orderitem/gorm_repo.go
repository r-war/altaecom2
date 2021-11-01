package orderitem

import (
	"AltaEcom/business/order"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type OrderItem struct {
	ID        int `gorm:"id;primaryKey;autoIncrement"`
	OrderID   int `gorm:"order_id;"`
	ProductID int `gorm:"product_id"`
	Qty       int `gorm:"qty"`
	Price     int `gorm:"price"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrderItemWithProductName struct {
	OrderItem
	ProductName string `gorm:"name"`
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}
func NewOrderItemTable(
	cartID int,
	item order.OrderItem,
) *OrderItem {
	return &OrderItem{
		ID:        item.ID,
		OrderID:   cartID,
		ProductID: item.ProductID,
		Qty:       item.Qty,
		Price:     item.Price,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
}

func (col *OrderItem) ToOrderItem() order.OrderItem {
	return order.OrderItem{
		ID:        col.ID,
		OrderID:   col.OrderID,
		ProductID: col.ProductID,
		Qty:       col.Qty,
		Price:     col.Price,
		CreatedAt: col.CreatedAt,
		UpdatedAt: col.UpdatedAt,
	}
}

func (r *Repository) AddItemToOrder(orderID int, item order.OrderItem) error {
	err := r.DB.Create(NewOrderItemTable(orderID, item)).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetOrderItemByUserID(OrderID int) (*[]order.OrderItem, error) {
	var itemsDetail *[]OrderItem

	err := r.DB.Raw("select t1.*, t2.name product_name from order_item t1 inner join products t2 on t2.id = t1.product_id where t1.order_id = ?", OrderID).Scan(itemsDetail).Error

	if err != nil {
		return nil, err
	}

	fmt.Println(itemsDetail)
	var result []order.OrderItem
	var temp order.OrderItem
	for _, val := range *itemsDetail {
		temp = val.ToOrderItem()
		//temp.ProductName = val.product.name
		result = append(result, temp)
	}
	return &result, nil
}

func (r *Repository) UpdateItemInOrder(orderID int, item order.OrderItem) error {
	var orderItem OrderItem

	err := r.DB.Find(&orderItem, "order_id = ? and id = ?", orderID, item.ID).Error
	// err := r.DB.Where("order_id = ? and id = ?", orderID, item.ID).First(&orderItem).Error
	if err != nil {
		return err
	}

	r.DB.Model(&orderItem).Updates(OrderItem{
		ID:        item.ID,
		ProductID: item.ProductID,
		Price:     item.Price,
		Qty:       item.Qty,
		UpdatedAt: time.Now(),
	})

	return nil
}

func (r *Repository) RemoveItemInOrder(orderID int, productID int) error {
	var orderItem OrderItem

	err := r.DB.Where("order_id = ? and product_id = ? ", orderID, productID).Find(&orderItem).Error

	if err != nil {
		return err
	}

	r.DB.Delete(&orderItem)
	return nil
}
