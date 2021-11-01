package main

import (
	"AltaEcom/api"
	"AltaEcom/config"
	"AltaEcom/modules/migration"

	authController "AltaEcom/api/auth"
	categoryController "AltaEcom/api/category"
	orderController "AltaEcom/api/order"
	productController "AltaEcom/api/product"
	transcationController "AltaEcom/api/transcation"
	userController "AltaEcom/api/user"

	adminService "AltaEcom/business/admin"
	authService "AltaEcom/business/auth"
	categoryService "AltaEcom/business/category"
	orderService "AltaEcom/business/order"
	productService "AltaEcom/business/product"
	transcationService "AltaEcom/business/transcation"
	userService "AltaEcom/business/user"

	adminRepository "AltaEcom/modules/admin"
	categoryRepository "AltaEcom/modules/category"
	orderRepo "AltaEcom/modules/order"
	orderItemRepo "AltaEcom/modules/orderitem"
	productRepository "AltaEcom/modules/product"
	transcationRepository "AltaEcom/modules/transcation"
	userRepository "AltaEcom/modules/user"

	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnection(conf *config.AppConfig) *gorm.DB {

	// config := map[string]string{
	// 	"DB_Username": "root",
	// 	"DB_Password": "",
	// 	"DB_Port":     "3306",
	// 	"DB_Host":     "127.0.0.1",
	// 	"DB_Name":     "altaecom",
	// }

	connectString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DbUsername,
		conf.DbPassword,
		conf.DbHost,
		conf.DbPort,
		conf.DbName)

	db, e := gorm.Open(mysql.Open(connectString), &gorm.Config{})

	if e != nil {
		panic(e)
	}

	migration.InitMigrate(db)

	return db

}

func main() {
	cfg := config.GetConfig()

	dbConnect := DBConnection(cfg)

	//initiate user repository
	userRepo := userRepository.NewGormDBRepository(dbConnect)

	//initiate user service
	userService := userService.NewService(userRepo)

	//initiate user controller
	userController := userController.NewController(userService)

	//admin
	adminRepo := adminRepository.NewGormDBRepository(dbConnect)

	adminService := adminService.NewService(adminRepo)

	//auth
	authService := authService.NewService(userService, adminService, adminRepo, userRepo, cfg)
	authController := authController.NewController(authService, cfg)

	//product
	productRepo := productRepository.NewRepository(dbConnect)

	productService := productService.NewService(adminService, productRepo)

	productController := productController.NewController(productService)

	//category
	categoryRepo := categoryRepository.NewGormDBRepository(dbConnect)

	categoryService := categoryService.NewService(adminService, categoryRepo)

	categoryController := categoryController.NewController(categoryService)

	//order
	orderRepo := orderRepo.NewRepository(dbConnect)
	orderItemRepo := orderItemRepo.NewRepository(dbConnect)

	orderService := orderService.NewService(orderRepo, orderItemRepo)

	orderController := orderController.NewController(orderService)

	transcationRepo := transcationRepository.NewRepository(dbConnect)
	transcationService := transcationService.NewService(transcationRepo, orderService)
	transcationController := transcationController.NewController(transcationService)

	e := echo.New()
	api.RegisterPath(
		e,
		userController,
		authController,
		productController,
		categoryController,
		orderController,
		transcationController,
		cfg,
	)

	func() {
		address := fmt.Sprintf(":%d", cfg.AppPort)

		if err := e.Start(address); err != nil {
			log.Info("Shutdown Echo Service")
		}

	}()
}
