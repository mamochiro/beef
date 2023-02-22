package test

import (
	"context"
	"github.com/mamochiro/beef/internals/config"
	"github.com/mamochiro/beef/internals/repository/rest/mocks"
	service "github.com/mamochiro/beef/internals/service/beef"
	"github.com/stretchr/testify/suite"
)

type PackageTestSuite struct {
	suite.Suite
	ctx      context.Context
	config   config.Configuration
	restRepo *mocks.Interface
	service  service.Interface
}

func (suite *PackageTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.config = config.NewConfiguration()
	suite.restRepo = &mocks.Interface{}
	suite.service = service.NewService(suite.config, suite.restRepo)
}
