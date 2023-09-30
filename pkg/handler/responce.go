package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type error struct {
	Message string `json:"message"`
}

func newErrorResponce(c *gin.Context, statusCode int, message string) { //*gin.Context-фреймворк для обработки HTTP-запросов и отправки ответов клиенту
	/*statusCode-статус код HTTP-ответа*/
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, error{message}) //выводит в формате json
}
