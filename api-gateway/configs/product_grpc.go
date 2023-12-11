package configs

// Deploy Config
import (
	"crypto/tls"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Deploy Cnfig
func ProductGrpc(target string) *grpc.ClientConn {
	creds := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.Dial(
		target, opts...,
	)
	if err != nil {
		log.Fatalf("Failed connect to grpc : %v", err)
	}

	return conn
}
