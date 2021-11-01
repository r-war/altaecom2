package response

import (
	"AltaEcom/business/product"
)

type GetProductResponse struct {
	ID          int    `json:"id"`
	CategoryID  int    `json:"category_id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Qty         int    `json:"qty"`
	Description string `json:"description"`
	Image       string `json:"image"`
	//UpdatedAt   time.Time `json:"updated_at`
}

func NewGetProductResponse(product product.Product) *GetProductResponse {
	var response GetProductResponse

	response.ID = product.ID
	response.CategoryID = product.CategoryId
	response.Name = product.Name
	response.Price = product.Price
	response.Qty = product.Qty
	response.Description = product.Description
	response.Image = product.Image
	//response.UpdatedAt = product.UpdatedAt

	return &response
}
