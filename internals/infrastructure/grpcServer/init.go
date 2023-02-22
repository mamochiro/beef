package grpcserver

import (
	grpcErrors "git.robodev.co/imp/shared-utility/grpc_errors"
	validatorUtils "git.robodev.co/imp/shared-utility/validator"
	"github.com/mamochiro/beef/internals/controller"
	"github.com/mamochiro/beef/internals/controller/beef"
	apiV1 "github.com/mamochiro/beef/pkg/api/v1"

	"github.com/mamochiro/beef/internals/config"
	"google.golang.org/grpc"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"
)

type Server struct {
	Config     config.Configuration
	Server     *grpc.Server
	HealthCtrl *controller.HealthZController
	BeefCtrl   *beef.Controller
}

// Configure ...
func (s *Server) Configure() {
	grpcHealthV1.RegisterHealthServer(s.Server, s.HealthCtrl)
	apiV1.RegisterBeefServiceServer(s.Server, s.BeefCtrl)
}

func NewServer(
	config config.Configuration,
	healthCtrl *controller.HealthZController,
	validator *validatorUtils.CustomValidator,
	beefCtrl *beef.Controller,
) *Server {
	option := grpc.ChainUnaryInterceptor(
		grpcErrors.UnaryServerInterceptor(),
		validatorUtils.UnaryServerInterceptor(validator),
	)

	s := &Server{
		Server:     grpc.NewServer(option, grpc.MaxRecvMsgSize(10*10e6), grpc.MaxSendMsgSize(10*10e6)),
		Config:     config,
		HealthCtrl: healthCtrl,
		BeefCtrl:   beefCtrl,
	}

	s.Configure()

	return s
}
