package request

import "AltaEcom/business/category"

type UpdateCategoryRequest struct {
	Name string `json:"name"`
}

func (update *UpdateCategoryRequest) ToCategorySpec() *category.CategorySpec {
	var spec category.CategorySpec

	spec.Name = update.Name

	return &spec
}
