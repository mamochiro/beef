package beef

import (
	"github.com/mamochiro/beef/internals/service/beef/wrapper"
)

type Controller struct {
	service wrapper.Wrapper
}

func NewController(
	service wrapper.Wrapper,
) *Controller {
	return &Controller{
		service: service,
	}
}
