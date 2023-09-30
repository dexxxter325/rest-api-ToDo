package main

import (
	todo "REST_API_ToDo"
	"REST_API_ToDo/pkg/handler"
	"REST_API_ToDo/pkg/repository"
	"REST_API_ToDo/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

// запуск программы
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter)) //задаем для логов(инфа о событиях в нашей приложухе) удобный json формат
	if err := initConfig(); err != nil {
		logrus.Fatalf("error intializizing configs:%s", err.Error())
	}
	if err := godotenv.Load(); err != nil { //godotenv.Load загружает переменные окружения из файла .env
		logrus.Fatalf("error loading env variables:%s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		//viper.GetString-для чтения строки (string) из конфигурационного файла или переменных окружения
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"), //os.Getenv() используется для получения значения переменной из файла .env
	})

	if err != nil {
		logrus.Fatalf("failed to initializate db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)                                                         //инициализируем сервер(чтобы потом можно было вызывать его методы)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil { //getstring-значение из вайпера по ключу
		logrus.Fatalf("error occured while running http server:%s", err.Error()) //err.error-наша красная ошибка
	}
}

/*Конфигурационный файл- файл,содержащий настройки и параметры для программного обеспечения или приложения*/
func initConfig() error {
	viper.AddConfigPath("configs") //имя нашей директории
	viper.SetConfigName("config")  //имя нашего файла
	return viper.ReadInConfig()    //считывает значения конфигов и записывает в вайпер
}
