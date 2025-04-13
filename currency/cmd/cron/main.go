package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/ArtyomYatsenko/currency/internal/clients/currency"
	"github.com/ArtyomYatsenko/currency/internal/config"
	"github.com/ArtyomYatsenko/currency/internal/database"
	"github.com/ArtyomYatsenko/currency/internal/migrations"
	"github.com/ArtyomYatsenko/currency/internal/repository"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"log"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	if err := run(); err != nil {
		log.Fatal(err)
	}

}

func run() error {

	logger, err := zap.NewProduction() // Создаю логер
	if err != nil {
		return fmt.Errorf("zap new profaction: %s", err)
	}
	defer logger.Sync()

	logger.Info("start...")

	configPath := flag.String("config", "./currency/configs", "path to the config file") // Получаю путь к конфигурации через параметры запуска
	flag.Parse()

	configApp, err := config.LoadConfig(*configPath) // Загружаю конфигурацию

	if err != nil {
		return fmt.Errorf("config load config: %s", err)
	}

	db, err := database.NewPostgresDB(configApp.DataBaseConfig) // Устанавливаю подключение к БД
	if err != nil {
		return fmt.Errorf("database new postgres db: %s", err)
	}

	currencyRepository := repository.NewCurrencyRepository(db) // Абстракция для запросов к БД

	migrator, err := migrations.NewMigrator(configApp.DataBaseConfig.DirMigrations, logger) // Создаю мигратор
	if err != nil {
		return fmt.Errorf("migrations new migrator %s", err)
	}

	err = migrator.ApplyMigrations(db) // Применяю миграции

	if err != nil {
		return fmt.Errorf("migrator apply migranions")
	}

	loc, err := time.LoadLocation("Europe/Moscow") // Создаю локацию так, как в контейнере другое время

	if err != nil {
		return fmt.Errorf("time load location %s", err)
	}

	client, err := currency.NewHttpClient(configApp.HttpClient, logger) // Создаю новый http клиент для подключения

	if err != nil {
		return fmt.Errorf("currenc new http client %s", err)
	}

	c := cron.New(cron.WithLocation(loc)) // Создаю крон планировщик

	specParam := fmt.Sprintf("%d %d * * *", configApp.TaskStartTime.Minute, configApp.TaskStartTime.Hour) // Указываю параметры выполнения задачи

	specParam = "*/1 * * * *" // УДАЛИТЬ!!!

	if _, err = c.AddFunc(specParam, func() { // Добавляю задачу в крон
		dailyTask(client, currencyRepository, logger)
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

func dailyTask(client *currency.Currency, currencyRepository *repository.CurrencyRepository, logger *zap.Logger) {

	data, err := client.FetchData()

	if err != nil {
		logger.Info("client fetch data", zap.Error(err))
		return
	}

	err = currencyRepository.AddCurrencies(data)
	if err != nil {
		logger.Info("currency repository add currency", zap.Error(err))
		return
	}

	log.Println(data)

}
