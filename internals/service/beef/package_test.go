package beef_test

import (
	"github.com/mamochiro/beef/internals/service/beef/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(test.PackageTestSuite))
}
