package request

import "AltaEcom/business/product"

type UpdateProductRequest struct {
	CategoryID  int    `json:"CategoryID"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Qty         int    `json:"qty"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func (update *UpdateProductRequest) ToProductSpec() *product.ProductSpec{
	var spec product.ProductSpec

	spec.Name = update.Name
	spec.CategoryID = update.CategoryID
	spec.Price = update.Price
	spec.Qty = update.Qty
	spec.Description = update.Description
	spec.Image = update.Image

	return &spec
}