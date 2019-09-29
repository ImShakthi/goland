package services

import (
	"github.com/imshakthi/goland/repositories"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	GetUser()
}

type userService struct {
	configService ConfigService
	userRepo      repositories.UserRepo
}

func NewUserService(configService ConfigService, userRepo repositories.UserRepo) UserService {
	return &userService{
		configService: configService,
		userRepo:      userRepo,
	}
}

func (service *userService) GetUser() {
	userDetail := service.userRepo.GetUserByID("1")
	logrus.Info("[INFO] user", userDetail)
}
