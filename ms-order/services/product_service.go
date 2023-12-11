package services

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"order/configs"
	pbProduct "order/internal/product"
	"order/middlewares"

	grpcMetadata "google.golang.org/grpc/metadata"
)

func CheckStock(ctx context.Context, req *pbProduct.CheckStockRequest) (*pbProduct.ListProductResponse, error) {
	// GRPC Auth
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	token, err := middlewares.SignJwtForGrpc()
	if err != nil {
		return nil, err
	}

	ctxWithAuth := grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)

	pHost := os.Getenv("PRODUCT_GRPC_HOST")
	pPort := os.Getenv("PRODUCT_GRPC_PORT")
	client := configs.ProductGrpc(fmt.Sprintf("%s:%s", pHost, pPort))

	productClient := pbProduct.NewProductServiceClient(client)

	product, err := productClient.CheckStock(ctxWithAuth, req)

	// product, err := productClient.CheckStock(ctx, req)
	if err != nil {
		log.Printf("Error from check stock service, err: %v\n", err)
		return nil, err
	}

	return product, nil
}
