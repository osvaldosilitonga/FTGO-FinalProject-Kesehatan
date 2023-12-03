package services

import (
	"context"
	"log"
	"os"

	"order/configs"
	pbProduct "order/internal/product"
)

func CheckStock(ctx context.Context, req *pbProduct.CheckStockRequest) (*pbProduct.ListProductResponse, error) {

	client := configs.ProductGrpc(os.Getenv("PRODUCT_GRPC_SERVER"))

	productClient := pbProduct.NewProductServiceClient(client)

	product, err := productClient.CheckStock(ctx, req)
	if err != nil {
		log.Printf("Error from check stock service, err: %v\n", err)
		return nil, err
	}

	return product, nil
}
