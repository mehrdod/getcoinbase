package main

import (
	"getcoinbase"
	"getcoinbase/configs"
	"getcoinbase/pkg/handler"
	"getcoinbase/pkg/repository"
	"getcoinbase/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}

func run() error {
	if err := configs.Init(); err != nil {
		return err
	}

	db, err := repository.NewMySqlDB(repository.Config{
		Host:     "localhost",
		Port:     "3306",
		Username: "root",
		Password: "qwerty",
		DBName:   "coinbase",
	})
	if err != nil {
		return err
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(getcoinbase.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		return err
	}
	return nil
}
