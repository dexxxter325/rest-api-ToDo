package handler

import (
	"REST_API_ToDo/pkg/service"
	"github.com/gin-gonic/gin"
)

// Handler /*отвечает за обработку событий или запросов*/
type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// InitRoutes /*Инициализация маршрутов(процесс определения и настройки маршрутов, по которым будут обрабатываться запросы от клиентов)*/
func (h *Handler) InitRoutes() *gin.Engine { //gin.engine-легкий фреймворк
	/*POST-создание(отправка данных на серв)
	GET -чтение(получение/чтение данных с сервера)
	PUT -обновление (отправка данных на сервер для обновления/модификации существующего ресурса)
	DELETE - удаление (удаление существующего ресурса на сервере)
	*/
	router := gin.New()           //роутер(маршрутизатор) для обработки входящих http запросов и опред.какие предпринимать действия с каждым маршрутом
	auth := router.Group("/auth") //маршруты,связанные с аунтефикацией
	{
		auth.POST("/sign-up", h.signUp) //регистрация
		auth.POST("/sign-in", h.signIn) //вход
	}
	api := router.Group("/api") //маршруты,связ. с api
	{
		lists := api.Group("/lists") //маршруты,связ.с операциями над списками
		{
			lists.POST("/", h.createList) //создается новый список при отправке на /...
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById) //для получения конкретного списка по его индефикатору.
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)
			items := lists.Group(":id/items") //маршруты,связанные с эл.списка в контексте опред.списка
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET(":item_id", h.getItemById)
				items.PUT(":item_id", h.updateItem)
				items.DELETE(":item_id", h.deleteItem)
			}
		}
	}
	return router

}
