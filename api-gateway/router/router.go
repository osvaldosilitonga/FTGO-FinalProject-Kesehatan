package router

import (
	"fmt"
	"gateway/configs"
	"gateway/controllers"
	"gateway/hooks"
	"gateway/middlewares"
	"gateway/service"
	"os"

	"github.com/labstack/echo/v4"
)

func Router(r *echo.Echo) {
	pHost := os.Getenv("PRODUCT_GRPC_HOST")
	pPort := os.Getenv("PRODUCT_GRPC_PORT")
	productGrpc := configs.ProductGrpc(fmt.Sprintf("%s:%s", pHost, pPort))

	oHost := os.Getenv("ORDER_GRPC_HOST")
	oPort := os.Getenv("ORDER_GRPC_PORT")
	orderGrpc := configs.ProductGrpc(fmt.Sprintf("%s:%s", oHost, oPort))

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
		user.GET("/profile/:id", userController.GetUserProfile)
	}

	order := v1.Group("/order")
	orderController := controllers.NewOrderController(orderService, paymentService, productService)
	order.Use(middlewares.RequireAuth)
	{
		// User Only
		order.POST("", orderController.CreateOrderProduct, middlewares.IsUser)
		order.POST("/cancel/:id", orderController.CancelOrder, middlewares.IsUser)

		// Owner and Admin/
		order.GET("/:id", orderController.OrderDetail)
		order.GET("/user/:id", orderController.FindByUserId)

		// Admin Only
		order.GET("/admin", orderController.ListOrder, middlewares.IsAdmin)
		order.PUT("/admin/confirm/:id", orderController.ConfirmOrder, middlewares.IsAdmin)
	}

	payment := v1.Group("/payment")
	paymentController := controllers.NewPaymentController(paymentService)
	payment.Use(middlewares.RequireAuth)
	{
		payment.GET("/:id", paymentController.FindByInvoiceID)
		payment.GET("/order/:id", paymentController.FindByOrderID)
		payment.GET("/user/:id", paymentController.FindByUserID)
	}

	xendit := v1.Group("/xendit")
	xenditHooks := hooks.NewXenditHooks(orderService, paymentService)
	{
		xendit.POST("/invoice", xenditHooks.InvoiceHooks)
	}

}
