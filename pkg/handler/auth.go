package handler

import (
	todo "REST_API_ToDo"
	"github.com/gin-gonic/gin"
	"net/http"
)

/* *gin.Context для обработки запросов и управления ответами*/

func (h *Handler) signUp(c *gin.Context) { //зарегаться,*gin.Context-набор методов и поля,урощающие обработку запросов и управление ответами.
	var input todo.User                        //хранение данных нового пользователя
	if err := c.BindJSON(&input); err != nil { //привязка данных JSON из запроса к переменной input
		newErrorResponce(c, http.StatusBadRequest, err.Error()) //StatusBadRequest-пользователь предоставил неккоректные данные в запросе
		return

	}
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error()) //внутренняя ошибка сервака
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{ //если все ок, то выводим id в json формате
		"id": id,
	})
}
func (h *Handler) signIn(c *gin.Context) { //вход

}
