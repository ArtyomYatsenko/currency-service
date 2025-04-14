package handler

import (
	"github.com/ArtyomYatsenko/gateway/internal/service"
	"github.com/gin-gonic/gin"
)

// Так как хендлеры должны иметь доступ к бизнес логики, добавлю доступ к ней
// Путем добавления в него поля структуры services

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/register", h.register)
		auth.POST("/login", h.login)
	}

	api := router.Group("/api")
	{
		api.GET("/datarate)", h.dateRate)

	}

	return router
}
