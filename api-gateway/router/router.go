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
	orderGrpc := configs.ProductGrpc(os.Getenv("ORDER_GRPC_SERVER"))

	redisClient := configs.InitRedis()

	productService := service.NewProductService(productGrpc)
	orderService := service.NewOrderService(orderGrpc)
	paymentService := service.NewPaymentService()
	userService := service.NewUserService()

	v1 := r.Group("/api/v1")

	product := v1.Group("/products")
	productController := controllers.NewProductController(productService)
	{
		product.GET("", productController.ListProduct)
		product.GET("/:id", productController.FindByID)

		// Admin Only
		product.POST("", productController.CreateProduct)
		product.PUT("/:id", productController.UpdateProduct)
		product.DELETE("/:id", productController.DeleteProduct)
	}

	user := v1.Group("/user")
	userController := controllers.NewUserController(userService, redisClient)
	{
		user.POST("/login", userController.Login)
		user.POST("/register", userController.Register)
	}

	order := v1.Group("/order")
	orderController := controllers.NewOrderController(orderService, paymentService)
	// order.Use(middlewares.RequireAuth)
	{
		order.POST("", orderController.CreateOrderProduct)
	}

	payment := v1.Group("/payment")
	paymentController := controllers.NewPaymentController()
	{
		payment.POST("", paymentController.Create)
	}

}
