package services

import (
	"github.com/BurntSushi/toml"
	"github.com/imshakthi/goland/models"
)

type ConfigService interface {
	GetDatabaseName() string
	GetDatabaseHost() string
	GetDatabaseUserName() string
	GetDatabasePassword() string
	GetDatabasePort() int
}

type configService struct {
	config models.Config
}

func NewConfigService() ConfigService {
	var config models.Config
	_, err := toml.DecodeFile("configurations/config.toml", &config)
	if err != nil {
		panic("Error while loading TOML data=" + err.Error())
	}
	return &configService{
		config: config,
	}
}

func (service *configService) GetRBMS() models.RBMS {
	return service.config.RBMS
}

func (service *configService) GetDatabase() models.Database {
	return service.GetRBMS().Database
}

func (service *configService) GetDatabaseUser() models.User {
	return service.GetRBMS().User
}

func (service *configService) GetDatabaseUserName() string {
	return service.GetDatabaseUser().Name
}
func (service *configService) GetDatabasePassword() string {
	return service.GetDatabaseUser().Password
}

func (service *configService) GetDatabaseName() string {
	return service.GetDatabase().Name
}

func (service *configService) GetDatabaseHost() string {
	return service.GetDatabase().Host
}

func (service *configService) GetDatabasePort() int {
	return service.GetDatabase().Port
}
