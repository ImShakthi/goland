package services

import (
	"fmt"
)

type ConfigService interface {
	GetConnectionURI() string
	//GetDatabaseName() string
	//GetDatabaseHost() string
	//GetDatabaseUserName() string
	//GetDatabasePassword() string
	//GetDatabasePort() int
}

type configService struct {
	//config models.Config
}

func NewConfigService() ConfigService {
	//var config models.Config
	//_, err := toml.DecodeFile("config/config.toml", &config)
	//if err != nil {
	//	panic("Error while loading TOML data=" + err.Error())
	//}
	return &configService{
		//config: config,
	}
}

func (s configService) GetConnectionURI() string {
	return fmt.Sprintf("host=%s port=5432 user=%s dbname=%s password=%s sslmode=disable",
		"ec2-54-246-85-151.eu-west-1.compute.amazonaws.com",
		"pphklugwpalshk",
		"d5ntof8uiq9jb9",
		"4422401351c2b73c2a3ba98f34abee7b27a840b4c2ae6f6c81d57db390afc276")

	//viper.GetString(config.DatabaseHost),
	//viper.GetString(config.DatabasePort),
	//viper.GetString(config.DatabaseUserName),
	//viper.GetString(config.DatabaseName),
	//viper.GetString(config.DatabasePassword))
}

//
//func (s configService) GetRBMS() models.RBMS {
//	return s.config.RBMS
//}
//
//func (s configService) GetDatabase() models.Database {
//	return s.GetRBMS().Database
//}
//
//func (s configService) GetDatabaseUser() models.User {
//	return s.GetRBMS().User
//}
//
//func (s configService) GetDatabaseUserName() string {
//	return s.GetDatabaseUser().Name
//}
//func (s configService) GetDatabasePassword() string {
//	return s.GetDatabaseUser().Password
//}
//
//func (s configService) GetDatabaseName() string {
//	return s.GetDatabase().Name
//}
//
//func (s configService) GetDatabaseHost() string {
//	return s.GetDatabase().Host
//}
//
//func (s configService) GetDatabasePort() int {
//	return s.GetDatabase().Port
//}
