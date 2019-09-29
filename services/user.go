package services

type UserService interface {
	GetUser()
}

type userService struct {
	configService ConfigService
}

func NewUserService(configService ConfigService) UserService {
	return &userService{
		configService: configService,
	}
}

func (service *userService) GetUser() {

}
