package request

import (
	"AltaEcom/business/product"
)

type InsertProductRequest struct {
	CategoryID  int    `json:"CategoryID"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Qty         int    `json:"qty"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func (ins InsertProductRequest) ToProductSpec() *product.ProductSpec {
	var spec product.ProductSpec

	spec.Name = ins.Name
	spec.CategoryID = ins.CategoryID
	spec.Price = ins.Price
	spec.Qty = ins.Qty
	spec.Description = ins.Description
	spec.Image = ins.Image

	return &spec
}