package test

import "errors"

func (suite *PackageTestSuite) TestGetBeefSummarySuccess() {
	resMock := []string{"Dog@@", "Test...!", "YoYo?"}
	suite.restRepo.Mock.On("BeefSummary", suite.ctx).Return(&resMock, nil).Once()
	summary, err := suite.service.BeefSummary(suite.ctx)

	expectedRes := map[string]int32{
		"dog":  1,
		"test": 1,
		"yoyo": 1,
	}
	suite.Equal(summary, expectedRes)
	suite.NoError(err)
	suite.NotNil(summary)
}

func (suite *PackageTestSuite) TestIdpListShouldThrowError() {
	suite.restRepo.Mock.On("BeefSummary", suite.ctx).Return(nil, errors.New("find resp error")).Once()
	summary, err := suite.service.BeefSummary(suite.ctx)
	suite.Error(err)
	suite.Nil(summary)
}
