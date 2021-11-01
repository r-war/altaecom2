package user

import (
	"AltaEcom/api/common"
	"AltaEcom/api/paginator"
	"AltaEcom/api/user/request"
	"AltaEcom/api/user/response"
	"AltaEcom/business/user"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

//Controller Get item API controller
type Controller struct {
	service user.Service
}

//NewController Construct item API controller
func NewController(service user.Service) *Controller {
	return &Controller{
		service,
	}
}

//GetItemByID Get item by ID echo handler
func (controller *Controller) FindUserByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := controller.service.FindUserByID(id)
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	response := response.NewGetUserResponse(*user)

	return c.JSON(common.NewSuccessResponse(response))
}

//FindAllUser Find All User with pagination handler
func (controller *Controller) FindAllUser(c echo.Context) error {

	pageQueryParam := c.QueryParam("page")
	rowPerPageQueryParam := c.QueryParam("row_per_page")

	skip, page, rowPerPage := paginator.CreatePagination(pageQueryParam, rowPerPageQueryParam)

	users, err := controller.service.FindAllUser(skip, rowPerPage)
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	response := response.NewGetAllUserResponse(users, page, rowPerPage)

	return c.JSON(common.NewSuccessResponse(response))
}

// InsertUser Create new user handler
func (controller *Controller) InsertUser(c echo.Context) error {
	insertUserRequest := new(request.InsertUserRequest)

	if err := c.Bind(insertUserRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := controller.service.InsertUser(*insertUserRequest.ToUpsertUserSpec(), "creator")
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}

// UpdateUser update existing user handler
func (controller *Controller) UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	updateUserRequest := new(request.UpdateUserRequest)

	if err := c.Bind(updateUserRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := controller.service.UpdateUser(id, updateUserRequest.Name, "modifier", updateUserRequest.Version)
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}
