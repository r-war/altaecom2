package category

import (
	"AltaEcom/api/category/request"
	"AltaEcom/api/category/response"
	"AltaEcom/api/common"
	"AltaEcom/business/category"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service category.Service
}

func NewController(service category.Service) *Controller {
	return &Controller{service}
}

func (c *Controller) GetCategories(ctx echo.Context) error {
	categories, err := c.service.GetCategories()
	if err != nil {
		return ctx.JSON(common.NewErrorBusinessResponse(err))
	}

	return ctx.JSON(common.NewSuccessResponse(
		response.GetCategories(categories).Categories,
	))
}

func (c *Controller) FindCategoryById(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	category, err := c.service.FindCategoryById(uint(id))
	if err != nil {
		return ctx.JSON(common.NewErrorBusinessResponse(err))
	}
	res := response.GetCategory(*category)

	return ctx.JSON(common.NewSuccessResponse(res))
}

func (c *Controller) InsertCategory(ctx echo.Context) error {
	var err error

	insertData := new(request.InsertCategoryRequest)

	if err = ctx.Bind(insertData); err != nil {
		return ctx.JSON(common.NewErrorBusinessResponse(err))
	}

	if err = c.service.InsertCategory(insertData.ToCategorySpec()); err != nil {
		return ctx.JSON(common.NewErrorBusinessResponse(err))
	}

	return ctx.JSON(common.NewSuccessResponseWithoutData())
}

func (c *Controller) UpdateCategory(ctx echo.Context) error {
	var err error

	id, _ := strconv.Atoi(ctx.Param("id"))

	updateCategory := new(request.UpdateCategoryRequest)
	if err = ctx.Bind(updateCategory); err != nil {
		return ctx.JSON(common.NewBadRequestResponse())
	}

	err = c.service.UpdateCategory(uint(id), updateCategory.ToCategorySpec())

	if err != nil {
		return ctx.JSON(common.NewErrorBusinessResponse(err))
	}

	return ctx.JSON(common.NewSuccessResponseWithoutData())
}

func (c *Controller) DeleteCategory(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	adminID, _ := strconv.Atoi(ctx.QueryParam("adminID"))

	if err := c.service.DeleteCategory(uint(id), adminID); err != nil {
		return ctx.JSON(common.NewErrorBusinessResponse(err))
	}
	return ctx.JSON(common.NewSuccessResponseWithoutData())

}
