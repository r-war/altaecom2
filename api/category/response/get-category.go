package response

import "AltaEcom/business/category"

type GetCategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func GetCategory(category category.Category) *GetCategoryResponse {
	var res GetCategoryResponse

	res.ID = category.ID
	res.Name = category.Name

	return &res
}
