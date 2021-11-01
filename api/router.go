package api

import (
	"AltaEcom/api/auth"
	"AltaEcom/api/category"
	"AltaEcom/api/middleware"
	"AltaEcom/api/order"
	"AltaEcom/api/product"
	"AltaEcom/api/transcation"
	"AltaEcom/api/user"
	"AltaEcom/config"

	"github.com/labstack/echo/v4"
)

func RegisterPath(
	e *echo.Echo,
	userController *user.Controller,
	authController *auth.Controller,
	productController *product.Controller,
	categoryController *category.Controller,
	orderController *order.Controller,
	transcationController *transcation.Controller,
	cfg *config.AppConfig,
) {
	if authController == nil || userController == nil {
		panic("Controller parameter cannot be nil")
	}

	//authentication with Versioning endpoint
	authV1 := e.Group("api/auth")
	authV1.POST("/login", authController.Login)
	authV1.POST("/register-admin", authController.RegisterAdmin)
	authV1.POST("/register-user", authController.RegisterUser)

	//user with Versioning endpoint
	userV1 := e.Group("api/users")
	userV1.Use(middleware.JWTMiddleware(*cfg))
	userV1.GET("/:id", userController.FindUserByID)
	userV1.GET("", userController.FindAllUser)
	userV1.POST("", userController.InsertUser)
	userV1.PUT("/:id", userController.UpdateUser)

	categoryV1 := e.Group("api/category")
	categoryV1.Use(middleware.JWTMiddleware(*cfg))
	categoryV1.GET("/:id", userController.FindUserByID)
	categoryV1.GET("", userController.FindAllUser)
	categoryV1.POST("", userController.InsertUser)
	categoryV1.PUT("/:id", userController.UpdateUser)

	productV1 := e.Group("api/product")
	productV1.Use(middleware.JWTMiddleware(*cfg))
	productV1.GET("/:id", userController.FindUserByID)
	productV1.GET("", userController.FindAllUser)
	productV1.POST("", userController.InsertUser)
	productV1.PUT("/:id", userController.UpdateUser)

	orderV1 := e.Group("api/cart")
	orderV1.Use(middleware.JWTMiddleware(*cfg))
	orderV1.GET("", orderController.GetOrderItemByUserID)
	orderV1.POST("", orderController.NewOrderByUserID)
	orderV1.POST("/add", orderController.AddItemToOrder)
	orderV1.POST("/update", orderController.UpdateItemInOrder)
	orderV1.DELETE("/delete/:productid", orderController.RemoveItemInOrder)

	transcationV1 := e.Group("api/checkout")
	transcationV1.Use(middleware.JWTMiddleware(*cfg))
	transcationV1.GET("", transcationController.GetAllTranscation)
	transcationV1.POST("", transcationController.CreateTranscation)
	transcationV1.POST("/status/:id", transcationController.UpdateTranscation)

	//health check
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})

}
