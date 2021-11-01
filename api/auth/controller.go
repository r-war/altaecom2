package auth

import (
	"AltaEcom/api/auth/request"
	"AltaEcom/api/auth/response"
	"AltaEcom/api/common"
	"AltaEcom/api/middleware"
	"AltaEcom/business/auth"
	"AltaEcom/config"
	utils "AltaEcom/util"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"

	echo "github.com/labstack/echo/v4"
)

//Controller Get item API controller
type Controller struct {
	service auth.Service
	cfg     config.AppConfig
}

//NewController Construct item API controller
func NewController(service auth.Service, cfg *config.AppConfig) *Controller {
	return &Controller{
		service,
		*cfg,
	}
}

//Login by given username and password will return JWT token
func (controller *Controller) Login(c echo.Context) error {
	loginRequest := new(request.LoginRequest)

	if err := c.Bind(loginRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	user, err := controller.service.Login(loginRequest.Username, loginRequest.IsAdmin)

	if !utils.CheckPasswordHash(loginRequest.Password, user.Password) {
		return c.JSON(common.NewErrorResponse(common.ControllerResponse{
			Code:    common.ErrForbidden,
			Message: "Password Not Match",
			Data:    nil,
		}))
	}

	claims := &middleware.JwtCustomClaims{
		Name:    user.Name,
		ID:      user.ID,
		Email:   user.Username,
		IsAdmin: loginRequest.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	generate, err := token.SignedString([]byte(controller.cfg.JWTConfig))
	if err != nil {
		return c.JSON(common.NewErrorResponse(common.ControllerResponse{
			Code:    common.ErrForbidden,
			Message: err.Error(),
			Data:    nil,
		}))
	}

	response := response.NewLoginResponse(generate)

	return c.JSON(common.NewSuccessResponse(response))
}

// InsertUser Create new user handler
func (controller *Controller) RegisterAdmin(c echo.Context) error {
	var insertUserRequest = new(request.RegisterAdminRequest)
	if err := c.Bind(insertUserRequest); err != nil {
		return c.JSON(common.NewErrorResponse(common.ControllerResponse{
			Code:    common.ErrBadRequest,
			Message: err.Error(),
			Data:    nil,
		}))
	}

	if insertUserRequest.Password != insertUserRequest.RePassword {
		return c.JSON(common.NewErrorResponse(common.ControllerResponse{
			Code:    common.ErrBadRequest,
			Message: "Password & Re type Password Not Match",
			Data:    nil,
		}))
	}
	admin, err := controller.service.RegisterAdmin(*insertUserRequest)
	if err != nil {
		return c.JSON(common.NewErrorResponse(common.ControllerResponse{
			Code:    common.ErrForbidden,
			Message: err.Error(),
			Data:    nil,
		}))
	}

	return c.JSON(common.NewSuccessResponse(admin))
}

func (controller *Controller) RegisterUser(c echo.Context) error {
	var insertUserRequest = new(request.RegisterUserRequest)

	if err := c.Bind(insertUserRequest); err != nil {
		return c.JSON(common.NewErrorResponse(common.ControllerResponse{
			Code:    common.ErrBadRequest,
			Message: err.Error(),
			Data:    nil,
		}))
	}

	if insertUserRequest.Password != insertUserRequest.RePassword {
		return c.JSON(common.NewErrorResponse(common.ControllerResponse{
			Code:    common.ErrBadRequest,
			Message: "Password & Re type Password Not Match",
			Data:    nil,
		}))
	}

	admin, err := controller.service.RegisterUser(*insertUserRequest)
	if err != nil {
		return c.JSON(common.NewErrorResponse(common.ControllerResponse{
			Code:    common.ErrForbidden,
			Message: err.Error(),
			Data:    nil,
		}))
	}

	return c.JSON(common.NewSuccessResponse(admin))
}

func (controller *Controller) Logout(c echo.Context) error {
	res := utils.Response{
		Code: http.StatusOK,
		Msg:  "Success Logout",
	}

	return c.JSON(http.StatusOK, res)
}
