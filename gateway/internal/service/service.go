package service

import (
	"github.com/ArtyomYatsenko/gateway/internal/dto"
	"github.com/ArtyomYatsenko/gateway/internal/repository"
)

// У меня два зоны отвественности (получения курса валют по датам "currency" и
// Работа с пользователями регистрация / авторизация
// Каждый интерфейс имеет свою зону ответственности, Service группирует их

type Authorization interface {
	CreateUser(user dto.User)
}

type Currency interface {
}

type Service struct {
	Authorization
	Currency
}

//Сервисы могут и будут обращаться к БД, поэтому передаю указатель на него у параметры

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
