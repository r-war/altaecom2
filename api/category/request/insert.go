package request

import "AltaEcom/business/category"

type InsertCategoryRequest struct {
	Name string `json:"name"`
}

func (ins InsertCategoryRequest) ToCategorySpec() *category.CategorySpec {
	var spec category.CategorySpec

	spec.Name = ins.Name

	return &spec
}
