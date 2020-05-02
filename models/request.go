package models

type UserDTO struct {
	Name string `json:"name"binding:"required"`
	Age  int    `json:"age"binding:"required"`
}
