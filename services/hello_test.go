package services_test

import (
	"github.com/golang/mock/gomock"
	"github.com/imshakthi/goland/models"
	"github.com/imshakthi/goland/services"
	"github.com/stretchr/testify/suite"

	"testing"
)

type HelloServiceTestSuite struct {
	suite.Suite
	mockCtrl     *gomock.Controller
	helloService services.HelloService
}

func TestNewHelloService(t *testing.T) {
	suite.Run(t, new(HelloServiceTestSuite))
}

func (suite *HelloServiceTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	config := loadConfig()
	suite.helloService = services.NewHelloService(config)
}

func loadConfig() models.Config {
	return models.Config{
		RBMS: models.RBMS{
			Database: models.Database{},
			User: models.User{
				Name: "Name",
			},
		},
	}
}

func (suite *HelloServiceTestSuite) TearDownTest() {
	suite.mockCtrl.Finish()
}

func (suite *HelloServiceTestSuite) TestHelloService_ShouldReturnHelloWorldWhenCalled() {
	message := suite.helloService.Hello()

	suite.Require().Equal("Name", message)
}
