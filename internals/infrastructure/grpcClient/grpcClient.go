package grpcclient

import (
	"log"

	"github.com/mamochiro/beef/internals/config"

	"google.golang.org/grpc"
)

// HTTPGRPCClient ...
type HTTPGRPCClient struct {
	Config config.Configuration
}

// Connect ...
func (client *HTTPGRPCClient) Connect() {
	// Set up a connection to the server.
	_, err := grpc.Dial("0.0.0.0:3000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	log.Println("Connect to service on", "0.0.0.0:3000")

}

// NewHTTPGRPCClient ...
func NewHTTPGRPCClient(config config.Configuration) *HTTPGRPCClient {
	// fmt.Printf("new grpc client: %+v\n", config)
	grpcClient := HTTPGRPCClient{
		Config: config,
	}

	grpcClient.Connect()

	return &grpcClient
}
