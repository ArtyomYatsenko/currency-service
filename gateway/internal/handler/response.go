package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

// Структура нужна для возврата ошибок, буду парсить ее в json

type error struct {
	Message string `json:"message"`
}

// Логируем и блокируем выполнение дальнейших обработчиков

func NewErrorResponse(c *gin.Context, statusCode int, message string) { //HELP
	log.Println(message)
	c.AbortWithStatusJSON(statusCode, error{message})
}
