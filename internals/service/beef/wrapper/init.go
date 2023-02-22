package wrapper

import (
	service "github.com/mamochiro/beef/internals/service/beef"
	"go.uber.org/dig"
)

type Wrapper struct {
	dig.In  `name:"wrapperBeef"`
	Service service.Interface
}
