package main

import (
	"flag"
	"fmt"
	"github.com/ArtyomYatsenko/gateway/internal/config"
	"github.com/ArtyomYatsenko/gateway/internal/handler"
	"github.com/ArtyomYatsenko/gateway/internal/repository"
	"github.com/ArtyomYatsenko/gateway/internal/server"
	"github.com/ArtyomYatsenko/gateway/internal/service"
	"go.uber.org/zap"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	logger, err := zap.NewProduction()

	if err != nil {
		return err
	}

	defer logger.Sync()

	logger.Info("start...")

	configPath := flag.String("config", "./gateway/configs", "path to the config file") // Получаем адрес конфигурации из параметров запуска
	flag.Parse()

	configApp, err := config.LoadConfig(*configPath)

	if err != nil {
		return err
	}

	logger.Info("config", zap.Any("", configApp.Server))

	logger.Info("conf rds", zap.Any("", configApp.DataBaseConfig))

	rdb := repository.NewRedisRepository(configApp.DataBaseConfig)
	fmt.Println(rdb)

	// Соблюдаю чистую архитектуру и реализую три слоя, handlers - транспортный
	// services - бизнес логика
	// repository - база данных
	base := repository.NewMyBase()
	repos := repository.NewRepository(base)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err = srv.Start(configApp.Server, handlers.InitRoutes()); err != nil {
		return err
	}

	return nil
}
