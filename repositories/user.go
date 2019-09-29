package repositories

import (
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
