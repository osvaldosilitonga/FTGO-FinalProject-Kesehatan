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

	productService := service.NewProductService(productGrpc)

	productController := controllers.NewProductController(productService)

	v1 := r.Group("/api/v1")

	product := v1.Group("/products")
	{
		product.GET("", productController.ListProduct)
		product.GET("/:id", productController.FindByID)

		// Admin Only
		product.POST("", productController.CreateProduct)
	}

}
