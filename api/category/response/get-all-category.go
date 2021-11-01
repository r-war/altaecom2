package response

import "AltaEcom/business/category"

type GetCategoriesResponse struct {
	Categories []GetCategoryResponse
}

func GetCategories(categories *[]category.Category) *GetCategoriesResponse {
	var res GetCategoriesResponse
	var category GetCategoryResponse

	for _, val := range *categories {
		category.ID = val.ID
		category.Name = val.Name

		res.Categories = append(res.Categories, category)
	}

	if res.Categories != nil {
		res.Categories = []GetCategoryResponse{}
	}

	return &res
}
