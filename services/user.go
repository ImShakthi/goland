package services

import (
	"errors"
	"github.com/imshakthi/goland/models"
	"github.com/imshakthi/goland/repositories"
)

type UserService interface {
	GetUser(name string) (models.UserResponse, error)
	GetUsers() ([]models.UserResponse, error)
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

func (service *userService) GetUser(id string) (models.UserResponse, error) {
	userDetail := service.userRepo.GetUser(id)
	if userDetail == repositories.EmptyUserModel {
		return models.UserResponse{}, errors.New("item not found")
	}
	return models.UserResponse{
		Name: userDetail.Name,
		Age:  userDetail.Age,
	}, nil
}

func (service *userService) GetUsers() ([]models.UserResponse, error) {
	return service.GetUsers()
}
