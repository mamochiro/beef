package controller

import (
	"context"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"
)

// HealthZController ...
type HealthZController struct{}

func (c *HealthZController) Check(ctx context.Context, request *grpcHealthV1.HealthCheckRequest) (*grpcHealthV1.HealthCheckResponse, error) {
	return &grpcHealthV1.HealthCheckResponse{
		Status: grpcHealthV1.HealthCheckResponse_SERVING,
	}, nil
}

func (c *HealthZController) Watch(request *grpcHealthV1.HealthCheckRequest, server grpcHealthV1.Health_WatchServer) error {
	//TODO implement me
	panic("implement me")
}

// NewHealthZController ...
func NewHealthZController() *HealthZController {
	return &HealthZController{}
}
