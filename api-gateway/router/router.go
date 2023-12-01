package router

import (
	"gateway/configs"
	"gateway/controllers"
	"gateway/service"
	"os"

	"github.com/labstack/echo/v4"
)

func Router(r *echo.Echo) {
	productGrpc := configs.ProductGrpc(os.Getenv("PRODUCT_GRPC_SERVER"))

	redisClient := configs.InitRedis()

	productService := service.NewProductService(productGrpc)
	userService := service.NewUserService()

	productController := controllers.NewProductController(productService)
	userController := controllers.NewUserController(userService, redisClient)

	v1 := r.Group("/api/v1")

	product := v1.Group("/products")
	{
		product.GET("", productController.ListProduct)
		product.GET("/:id", productController.FindByID)

		// Admin Only
		product.POST("", productController.CreateProduct)
	}

	user := v1.Group("/user")
	{
		user.POST("/login", userController.Login)
	}

}
