package transcation

import (
	"AltaEcom/api/common"
	"AltaEcom/api/middleware"
	"AltaEcom/api/transcation/request"
	"AltaEcom/api/transcation/response"
	"AltaEcom/business/transcation"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service transcation.Service
}

func NewController(service transcation.Service) *Controller {
	return &Controller{service}
}

func (controller *Controller) GetAllTranscation(c echo.Context) error {
	transcations, err := controller.service.GetAllTranscation()
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponse(
		response.GetTranscationsResponse(transcations).Transcation,
	))
}

func (controller *Controller) GetTranscationById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transcation, err := controller.service.GetTranscationById(id)

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	res := response.NewTranscationResponse(*transcation)

	return c.JSON(common.NewSuccessResponse(res))
}

func (controller *Controller) CreateTranscation(c echo.Context) error {
	//var err error
	claims := middleware.GetTokenFromContext(c)

	var transcation = new(request.TranscationRequest)

	if err := c.Bind(transcation); err != nil {
		return c.JSON((common.NewErrorBusinessResponse(err)))
	}

	if err := controller.service.CreateTranscation(claims.ID, transcation.ToTranscationSpec()); err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}

func (controller *Controller) UpdateTranscation(c echo.Context) error {
	var err error
	id, _ := strconv.Atoi(c.Param("id"))

	transcation := new(request.TranscationRequest)
	if err = c.Bind(transcation); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err = controller.service.UpdateTranscation(id, transcation.ToTranscationSpec())
	if err != nil {
		return c.JSON((common.NewErrorBusinessResponse(err)))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}
