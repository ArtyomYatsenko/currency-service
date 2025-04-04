package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ArtyomYatsenko/currency/internal/config"
	"github.com/ArtyomYatsenko/currency/internal/database"
	"github.com/robfig/cron/v3"
	"io"
	"log"
	"net/http"
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

	configApp, err := config.LoadConfig()

	if err != nil {
		return fmt.Errorf("config load config: %s", err)
	}

	db, err := database.NewPostgresDB(configApp.DataBaseConfig)

	if err != nil {
		return fmt.Errorf("database new postgres db: %s", err)
	}

	c := cron.New()

	specParam := fmt.Sprintf("%d %d * * *", configApp.TaskStartTime.Minute, configApp.TaskStartTime.Hour)

	fmt.Println(configApp)

	if _, err = c.AddFunc(specParam, dailyTask); err != nil {
		return fmt.Errorf("cron add func: %s", err)

	}

	c.Start()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer cancel()

	<-ctx.Done()

	c.Stop()

	if err = db.Close(); err != nil {
		log.Printf("database close: %s", err)
	}

	return nil

}

func dailyTask() {

	log.Println("dailiTask")
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get("https://latest.currency-api.pages.dev/v1/currencies/rub.json")
	if err != nil {
		log.Printf("http client get: %s", err)
	}

	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Printf("http resp body close: %s", err)
		}
	}()

	bodyBytes, err := io.ReadAll(resp.Body)

	var data map[string]interface{}

	if err = json.Unmarshal(bodyBytes, &data); err != nil {
		log.Printf("json unmarshal: %s", err)
	}

}
