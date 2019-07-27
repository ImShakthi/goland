package services

import "goland/constants"

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
