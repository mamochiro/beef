package container

import (
	"fmt"
	"github.com/mamochiro/beef/internals/infrastructure/httpClient"
	"github.com/mamochiro/beef/internals/repository/rest"
	"github.com/mamochiro/beef/internals/service/beef"
	"net/http"

	"git.robodev.co/imp/shared-utility/validator"
	"github.com/mamochiro/beef/internals/config"
	"github.com/mamochiro/beef/internals/controller"
	beefCtrl "github.com/mamochiro/beef/internals/controller/beef"
	"github.com/mamochiro/beef/internals/infrastructure/database"
	grpcServer "github.com/mamochiro/beef/internals/infrastructure/grpcServer"
	httpServer "github.com/mamochiro/beef/internals/infrastructure/httpServer"
	"github.com/mamochiro/beef/internals/infrastructure/jaeger"
	"github.com/mamochiro/beef/internals/infrastructure/logrus"
	"github.com/mamochiro/beef/internals/repository/postgres"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/mamochiro/beef/internals/utils"
	"go.uber.org/dig"
)

type Container struct {
	container *dig.Container
}

func (c *Container) Configure() error {
	servicesConstructors := []interface{}{
		config.NewConfiguration,
		grpcServer.NewServer,
		//database.NewServerBase,
		http.NewServeMux,
		httpServer.NewServer,
		runtime.NewServeMux,
		jaeger.NewJaeger,
		logrus.NewLog,
		controller.NewHealthZController,
		validator.NewCustomValidator,
		postgres.NewRepository,
		utils.NewUtils,
		utils.NewCustomValidator,
		beefCtrl.NewController,
		beef.NewService,
		httpClient.NewHttpClient,
		rest.NewClient,
	}

	for _, service := range servicesConstructors {
		if err := c.container.Provide(service); err != nil {
			return err
		}
	}
	appConfig := config.NewConfiguration()
	jaeger.NewJaeger(appConfig)
	return nil
}

func (c *Container) Start() error {
	fmt.Println("Start Container")

	if err := c.container.Invoke(func(s *grpcServer.Server, h *httpServer.Server, v *validator.CustomValidator) {
		go func() {
			_ = h.Start()
		}()
		s.Start()

	}); err != nil {
		fmt.Printf("%s", err)

		return err
	}

	return nil
}

// MigrateDB ...
func (c *Container) MigrateDB() error {
	fmt.Println("Start Container DB")

	if err := c.container.Invoke(func(d *database.DB) {
		d.MigrateDB()
	}); err != nil {
		return err
	}

	return nil
}

func NewContainer() (*Container, error) {
	d := dig.New()

	container := &Container{
		container: d,
	}

	if err := container.Configure(); err != nil {
		return nil, err
	}

	return container, nil
}
