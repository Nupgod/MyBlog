package main

import (
	"ex01"
	"ex01/pkg/handler"
	"ex01/pkg/repository"
	"ex01/pkg/services"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("Config initialization error: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBname:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("Faild to init DB: %s", err)
	}

	repo := repository.NewRepository(db)
	serv := services.NewService(repo)
	handlers := handler.NewHandler(serv)

	srv := new(ex01.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Run server error: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
