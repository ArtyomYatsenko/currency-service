package main

import (
	"context"
	"embed"
	"fmt"
	"github.com/ArtyomYatsenko/currency/internal/clients/currency"
	"github.com/ArtyomYatsenko/currency/internal/config"
	"github.com/ArtyomYatsenko/currency/internal/database"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//go:embed currency/internal/migrations/*.sql
var MigrationsFS embed.FS

func main() {

	if err := run(); err != nil {
		log.Fatal(err)
	}

}

func run() error {

	logger, err := zap.NewProduction() // Создаю логер

	logger.Info("start...")
	if err != nil {
		return fmt.Errorf("zap new profaction: %s", err)
	}
	defer logger.Sync()

	configPath := os.Getenv("CONFIG_PATH") //Читаю переменные путь к конфигурации из переменной окружения

	if configPath == "" {
		configPath = "currency/configs" // Указываем путь по умолчанию
	}

	configApp, err := config.LoadConfig(configPath) // Загружаю конфигурацию

	if err != nil {
		return fmt.Errorf("config load config: %s", err)
	}

	db, err := database.NewPostgresDB(configApp.DataBaseConfig) // Устанавливаю подключение к БД

	if err != nil {
		return fmt.Errorf("database new postgres db: %s", err)
	}

	loc, err := time.LoadLocation("Europe/Moscow") // Создаю локацию так, как в контейнере другое время

	if err != nil {
		return fmt.Errorf("time load location %s", err)
	}

	client, err := currency.NewHttpClient(configApp.HttpClient.Timeout, logger) // Создаю новый http клиент для подключения

	if err != nil {
		return fmt.Errorf("currenc new http client %s", err)
	}

	c := cron.New(cron.WithLocation(loc)) // Создаю крон планировщик

	specParam := fmt.Sprintf("%d %d * * *", configApp.TaskStartTime.Minute, configApp.TaskStartTime.Hour) // Указываю параметры выполнения задачи

	specParam = "*/1 * * * *" // УДАЛИТЬ!!!

	if _, err = c.AddFunc(specParam, func() { // Добавляю задачу в крон
		dailyTask(client)
	}); err != nil {
		return fmt.Errorf("cron add func: %s", err)
	}

	c.Start() // Запуск крона в отдельной горутине

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT) // graceful shutdown

	defer cancel()

	<-ctx.Done()

	c.Stop()

	if err = db.Close(); err != nil {
		return fmt.Errorf("database close: %s", err)
	}

	return nil

}

func dailyTask(client *currency.Currency) {
	data, err := client.FetchData()

	if err != nil {
		log.Printf("client fetch data: %s", err)
		return
	}

	log.Println(data)
}
