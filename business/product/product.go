package product

import "time"

type Product struct {
	ID           int
	CategoryId   int
	CategoryName string
	Name         string
	Price        int
	Qty          int
	Description  string
	Image        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

func NewProduct(
	CategoryId int,
	Name string,
	Price int,
	Qty int,
	Description string,
	Image string,
	CreatedAt time.Time,
	UpdatedAt time.Time) Product {

	return Product{
		CategoryId:  CategoryId,
		Name:        Name,
		Price:       Price,
		Qty:         Qty,
		Description: Description,
		Image:       Image,
		CreatedAt:   CreatedAt,
		UpdatedAt:   UpdatedAt,
	}
}

func ModifyProduct(
	name string,
	price int,
	categoryId int,
	qty int,
	description string,
	image string,
	updateAt time.Time,
) Product {
	return Product{
		Name:        name,
		Price:       price,
		CategoryId:  categoryId,
		Qty:         qty,
		Description: description,
		Image:       image,
		UpdatedAt:   updateAt,
	}
}
