package response

import "AltaEcom/business/product"

type GetProductsResponse struct {
	Products []GetProductResponse `json:"products"`
}

func NewGetProductsResponse(products []product.Product) GetProductsResponse {

	getProductsResponse := GetProductsResponse{}
	var productsResponse GetProductResponse

	for _, val := range products {

		productsResponse.ID = val.ID
		productsResponse.CategoryID = val.CategoryId
		productsResponse.Name = val.Name
		productsResponse.Price = val.Price
		productsResponse.Qty = val.Qty
		productsResponse.Description = val.Description
		productsResponse.Image = val.Image

		getProductsResponse.Products = append(getProductsResponse.Products, productsResponse)
	}

	if getProductsResponse.Products == nil {
		getProductsResponse.Products = []GetProductResponse{}
	}

	return getProductsResponse
}
