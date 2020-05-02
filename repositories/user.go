package repositories

import (
	"fmt"
	"github.com/imshakthi/goland/models"
	"github.com/jinzhu/gorm"
)

const userTableName = "user_details"

var (
	//id = "id"
	//name = "name"
	//age  = "age"
	EmptyUserModel = models.UserDetail{}
)

type UserRepo interface {
	GetUser(userID string) models.UserDetail
	GetUsers() []models.UserDetail
	AddUser(user models.UserDTO) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}
func (r *userRepo) GetUser(userID string) models.UserDetail {
	userDetail := models.UserDetail{}
	r.db.Table(userTableName).
		Where("id = ? ", userID).
		Find(&userDetail)
	return userDetail
}

func (r *userRepo) GetUsers() []models.UserDetail {
	var users []models.UserDetail

	r.db.Table(userTableName).
		Find(&users)

	return users
}

func (r *userRepo) AddUser(user models.UserDTO) error {
	affected := r.db.Table(userTableName).Create(user).RowsAffected

	if affected == 0 {
		return fmt.Errorf("unable to add user to database")
	}
	return nil
}
