package beef

import (
	"github.com/mamochiro/beef/internals/config"
	"github.com/mamochiro/beef/internals/repository/rest"
)

type Service struct {
	config   config.Configuration
	restRepo rest.Interface
}

func NewService(
	c config.Configuration,
	rest rest.Interface,
) (service Interface) {
	return Service{
		config:   c,
		restRepo: rest,
	}
}
