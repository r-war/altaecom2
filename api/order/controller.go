package order

import (
	"AltaEcom/api/common"
	"AltaEcom/api/middleware"
	"AltaEcom/api/order/request"
	"AltaEcom/api/order/response"
	"AltaEcom/business/order"
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service order.Service
}

func NewController(service order.Service) *Controller {
	return &Controller{service}
}

func (controller *Controller) GetOrderByUserID(c echo.Context) error {
	claims := middleware.GetTokenFromContext(c)

	order, err := controller.service.GetOrderByUserID(claims.ID)

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	res := response.NewGetOrderResponse(order)

	return c.JSON(common.NewSuccessResponse(res))
}

func (controller *Controller) NewOrderByUserID(c echo.Context) error {
	//var err error
	claims := middleware.GetTokenFromContext(c)

	addOrder := new(request.OrderRequest)

	if err := c.Bind(addOrder); err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}
	addOrder.UserID = claims.ID

	order, err := controller.service.NewOrderByUserID(addOrder.UserID)

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}
	response := response.NewGetOrderResponse(order)

	return c.JSON(common.NewSuccessResponse(response))
}

func (controller *Controller) GetOrderItemByUserID(c echo.Context) error {
	claims := middleware.GetTokenFromContext(c)

	orderItem, err := controller.service.GetOrderItemByUserID(claims.ID)
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}
	return c.JSON(common.NewSuccessResponse(orderItem))
}

func (controller *Controller) AddItemToOrder(c echo.Context) error {
	var item = new(request.OrderItemRequest)

	claims := middleware.GetTokenFromContext(c)
	order, _ := controller.service.GetOrderByUserID(claims.ID)

	if err := c.Bind(item); err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	fmt.Println(item.ToProductSpec())

	err := controller.service.AddItemToOrder(order.ID, item.ToProductSpec())
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}
	return c.JSON(common.NewSuccessResponseWithoutData())
}

func (controller *Controller) UpdateItemInOrder(c echo.Context) error {
	var item = new(request.OrderItemRequest)

	if err := c.Bind(item); err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	claims := middleware.GetTokenFromContext(c)

	order, _ := controller.service.GetOrderByUserID(claims.ID)

	// fmt.Println(item.ToProductSpec())

	err := controller.service.UpdateItemInOrder(order.ID, item.ToProductSpec())

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}

func (controller *Controller) RemoveItemInOrder(c echo.Context) error {
	claims := middleware.GetTokenFromContext(c)

	order, err1 := controller.service.GetOrderByUserID(claims.ID)
	productID, err2 := strconv.Atoi(c.Param("productid"))

	if err1 != nil || err2 != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := controller.service.RemoveItemInOrder(order.ID, productID)
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}
	return c.JSON(common.NewSuccessResponseWithoutData())

}
