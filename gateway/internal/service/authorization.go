package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/ArtyomYatsenko/gateway/internal/dto"
	"github.com/ArtyomYatsenko/gateway/internal/repository"
)

const salt = "eqrfasgrfdfshsfsdfa"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// Передаем структуру еще на слой ниже в репозиторий

func (s *AuthService) CreateUser(user dto.User) {
	fmt.Println(user)
	user.Password = generatePasswordHash(user.Password)
	fmt.Println(user)
	s.repo.CreateUser(user)
}

// Типо хранить нельзя в открытом виде и тд, хеширую
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
