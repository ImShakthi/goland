package repositories

import (
	"github.com/imshakthi/goland/models"
	"github.com/jinzhu/gorm"
)

const UserTableName = "user_details"

var (
	ID   = "id"
	Name = "name"
	Age  = "age"
)

type UserRepo interface {
	GetUserByID(userID string) models.UserDetail
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}
func (r *userRepo) GetUserByID(userID string) models.UserDetail {
	userDetail := models.UserDetail{}
	r.db.Table(UserTableName).Where(ID, " = ? ", userID).Find(&userDetail)
	return userDetail
}
