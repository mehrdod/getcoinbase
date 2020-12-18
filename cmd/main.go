package main

import (
	"getcoinbase"
	"getcoinbase/configs"
	"getcoinbase/pkg/handler"
	"getcoinbase/pkg/repository"
	"getcoinbase/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
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

	if err := godotenv.Load(); err != nil {
		return err
	}

	db, err := repository.NewMySqlDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASS"),
		DBName:   viper.GetString("db.name"),
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
