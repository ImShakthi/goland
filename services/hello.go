package services

import (
	"github.com/imshakthi/goland/models"
)

type HelloService interface {
	Hello() string
}

type helloService struct {
	config models.Config
}

func NewHelloService(config models.Config) HelloService {
	return &helloService{
		config: config,
	}
}

func (service *helloService) Hello() string {
	return service.config.RBMS.User.Name
}
