package configs

// Deploy Config
import (
	// "crypto/tls"
	"log"

	"google.golang.org/grpc"
	// "google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

// Deploy Cnfig
func ProductGrpc(target string) *grpc.ClientConn {
	// creds := credentials.NewTLS(&tls.Config{
	// 	InsecureSkipVerify: true,
	// })

	opts := []grpc.DialOption{
		// grpc.WithUnaryInterceptor(Interceptor),
		// grpc.WithTransportCredentials(creds),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// target := os.Getenv("PRODUCT_GRPC_SERVER")
	conn, err := grpc.Dial(
		target, opts...,
	)
	if err != nil {
		log.Fatalf("Failed connect to grpc : %v", err)
	}

	return conn
}

// Local Config
// import (
// 	// "crypto/tls"
// 	"log"

// 	"google.golang.org/grpc"
// 	// "google.golang.org/grpc/credentials"
// 	"google.golang.org/grpc/credentials/insecure"
// )

// Local Config
// func ProductGrpc(target string) *grpc.ClientConn {
// 	// creds := credentials.NewTLS(&tls.Config{
// 	// 	InsecureSkipVerify: true,
// 	// })

// 	opts := []grpc.DialOption{
// 		// grpc.WithUnaryInterceptor(Interceptor),
// 		// grpc.WithTransportCredentials(creds),
// 		grpc.WithTransportCredentials(insecure.NewCredentials()),
// 	}

// 	// target := os.Getenv("PRODUCT_GRPC_SERVER")
// 	conn, err := grpc.Dial(
// 		target, opts...,
// 	)
// 	if err != nil {
// 		log.Fatalf("Failed connect to grpc : %v", err)
// 	}

// 	return conn
// }
