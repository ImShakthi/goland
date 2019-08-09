package models

type UserRequest struct {
	Name    string `json:"name" binding:"required"`
	Checked bool   `json:"checked" binding:"required"`
}
