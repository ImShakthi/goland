package services

import (
	"errors"
	"github.com/imshakthi/goland/models"
	"github.com/imshakthi/goland/repositories"
)

type UserService interface {
	GetUser(name string) (models.UserDTO, error)
	GetUsers() ([]models.UserDTO, error)
	AddUser(user models.UserDTO) error
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

func (s *userService) GetUser(id string) (models.UserDTO, error) {
	userDetail := s.userRepo.GetUser(id)
	if userDetail == repositories.EmptyUserModel {
		return models.UserDTO{}, errors.New("item not found")
	}
	return models.UserDTO{
		Name: userDetail.Name,
		Age:  userDetail.Age,
	}, nil
}

func (s *userService) GetUsers() ([]models.UserDTO, error) {
	return s.GetUsers()
}

func (s *userService) AddUser(user models.UserDTO) error {
	return s.userRepo.AddUser(user)
}
