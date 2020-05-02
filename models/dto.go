package models

type UserDetail struct {
	ID   string
	Name string
	Age  int
}

func (u UserDetail) MapToUserDTO() UserDTO {
	return UserDTO{
		Name: u.Name,
		Age:  u.Age,
	}
}
