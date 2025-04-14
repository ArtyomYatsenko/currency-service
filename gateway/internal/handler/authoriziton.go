package handler

import (
	"github.com/ArtyomYatsenko/gateway/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) register(c *gin.Context) {

	var input dto.User

	// Преобразовываю json и проверяю валидацию
	if err := c.BindJSON(&input); err != nil {
		// Если косяк возвращаю 400 (предоставлены некорректные данные в запросе и текст ошибки)
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//Передаю данные на слой ниже

	h.services.Authorization.CreateUser(input)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": 1,
	})

}

func (h *Handler) login(c *gin.Context) {

}
