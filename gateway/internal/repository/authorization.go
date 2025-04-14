package repository

import (
	"fmt"
	"github.com/ArtyomYatsenko/gateway/internal/dto"
)

type AuthRepository struct {
	db MyBase
}

func NewAuthRepository(db MyBase) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user dto.User) {
	r.db[user.Username] = r.db[user.Password]
	fmt.Println(r.db)
}
