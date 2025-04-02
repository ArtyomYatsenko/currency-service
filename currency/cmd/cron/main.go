package main

import (
	"fmt"
	"github.com/ArtyomYatsenko/currency/internal/config"
	"github.com/ArtyomYatsenko/currency/internal/db"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"log"
)

func main() {

	if err := config.LoadConfig(); err != nil {
		log.Fatalf("error loading configuration: %s", err.Error())
	}

	_, err := db.NewPostgresDB(db.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("password"),
		DBName:   viper.GetString("dbname"),
		SSLMode:  viper.GetString("sslmode"),
	})

	if err != nil {
		log.Fatalf("initialization error db: %s", err.Error())
	}

	c := cron.New()

	specParam := fmt.Sprintf("%d %d * * *", viper.GetInt("task_start_time.minute"), viper.GetInt("task_start_time.hour"))

	if _, err := c.AddFunc(specParam, dailyTask); err != nil {
		log.Fatalf("error in starting cron task: %s", err.Error())
	}

	c.Start()

	defer c.Stop()

	select {}

}

func dailyTask() {
	fmt.Println("Выполнение задачи")
}
