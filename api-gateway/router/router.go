package router

import (
	"gateway/configs"
	"gateway/controllers"
	"gateway/middlewares"
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
		product.POST("", productController.CreateProduct, middlewares.RequireAuth, middlewares.IsAdmin)
		product.PUT("/:id", productController.UpdateProduct, middlewares.RequireAuth, middlewares.IsAdmin)
		product.DELETE("/:id", productController.DeleteProduct, middlewares.RequireAuth, middlewares.IsAdmin)
	}

	user := v1.Group("/user")
	userController := controllers.NewUserController(userService, redisClient)
	{
		user.POST("/login", userController.Login)
		user.POST("/register", userController.Register)
		user.POST("/register/admin", userController.RegisterAdmin)

		// User and Admin
		user.PUT("/profile/:id", userController.UpdateUserProfile, middlewares.RequireAuth)
		user.GET("/profile/:id", userController.GetUserProfile, middlewares.RequireAuth)
	}

	order := v1.Group("/order")
	orderController := controllers.NewOrderController(orderService, paymentService, productService)
	order.Use(middlewares.RequireAuth)
	{
		// User Only
		order.POST("", orderController.CreateOrderProduct, middlewares.IsUser)
		order.POST("/cancel/:id", orderController.CancelOrder, middlewares.IsUser)

		// Admin Only
		order.GET("/admin", orderController.ListOrder, middlewares.IsAdmin)
		// order.GET("/:id", orderController.FindByID, middlewares.IsAdmin)
		// order.PUT("/confirm/:id", orderController.UpdateOrder, middlewares.IsAdmin)
	}

	// payment := v1.Group("/payment")
	// paymentController := controllers.NewPaymentController()
	// {
	// 	payment.POST("", paymentController.Create)
	// }

}
