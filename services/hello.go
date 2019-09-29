package services

import (
	"github.com/imshakthi/goland/constants"
)

type HelloService interface {
	Hello() string
}

type helloService struct {
}

func NewHelloService() HelloService {
	return &helloService{}
}

func (service *helloService) Hello() string {
	return constants.HelloWorld
}
