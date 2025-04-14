package dto

type User struct {
	Id       int    `json:"-"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
