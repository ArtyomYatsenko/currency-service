package repository

import "github.com/ArtyomYatsenko/gateway/internal/dto"

// Здесь у меня одна зона ответсвенности, только работа с пользователями, т.к. только их пишем в БД
// О других действиях репозиторую не нужно знать

type Authorization interface {
	CreateUser(user dto.User)
}

type Repository struct {
	Authorization
}

func NewRepository(myBase MyBase) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(myBase),
	}
}
