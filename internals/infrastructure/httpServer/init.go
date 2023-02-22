package httpServer

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/mamochiro/beef/internals/config"
	"github.com/mamochiro/beef/internals/controller/beef"
	apiV1 "github.com/mamochiro/beef/pkg/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
	"strconv"
)

type Server struct {
	Config   config.Configuration
	Server   *runtime.ServeMux
	HttpMux  *http.ServeMux
	BeefCtrl *beef.Controller
}

func (s *Server) Configure(ctx context.Context, opts []grpc.DialOption) {
	apiV1.RegisterBeefServiceHandlerFromEndpoint(ctx, s.Server, "0.0.0.0:"+strconv.Itoa(s.Config.Port), opts)
}

func NewServer(
	config config.Configuration,
	rmux *runtime.ServeMux,
	httpMux *http.ServeMux,
	beefCtrl *beef.Controller,
) *Server {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	s := &Server{
		Config:   config,
		Server:   setRuntimeConfig(),
		HttpMux:  httpMux,
		BeefCtrl: beefCtrl,
	}

	s.Configure(context.Background(), opts)
	return s
}

func setRuntimeConfig() *runtime.ServeMux {
	return runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
			Marshaler: &runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseProtoNames:   true,
					EmitUnpopulated: false,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{
					DiscardUnknown: true,
				},
			},
		}),
	)
}
