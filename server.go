package todo

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server //для запуска http серв.
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,       /*адресс серв*/
		Handler:        handler,          //обработчик запросов
		MaxHeaderBytes: 1 << 20,          //макс.размер заголовка в байтах
		ReadTimeout:    10 * time.Second, //время ожидания серв. для чтения
		WriteTimeout:   10 * time.Second, //время ожидания серв. для записи
	}
	return s.httpServer.ListenAndServe() //обрабатывает входящие запросы от клиента
}
func (s *Server) Shutdown(ctx context.Context) error { //останавливаем наш сервер(в конце)
	//ctx-макс. время ожидания завершения соединений перед закрытиям серв.
	//shutdown-серв.перестает принимать новые запросы и начинает закрывать соединения
	return s.httpServer.Shutdown(ctx)
}
