package main

import (
	"flag"
	"github.com/ArtyomYatsenko/gateway/internal/config"
	"github.com/ArtyomYatsenko/gateway/internal/server"
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

	srv := &server.Server{}

	if err = srv.Start(configApp.Server); err != nil {
		return err
	}

	return nil
}
