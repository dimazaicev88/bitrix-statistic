package main

import (
	"bitrix-statistic/internal/app"
	"bitrix-statistic/internal/config"
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/services"
	"bitrix-statistic/internal/storage"
	"bitrix-statistic/internal/tasks"
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/hibiken/asynq"
	"github.com/huandu/go-sqlbuilder"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

//TODO для  всех методов добавить по возможности одинаковые ошибки.

func main() {
	sqlbuilder.DefaultFlavor = sqlbuilder.ClickHouse
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file", err.Error())
	}
	cfg := config.GetServerConfig()

	fb := fiber.New()
	chClient, err := storage.NewClickHouseClient(config.GetServerConfig())
	if err != nil {
		logrus.Fatal(err)
	}

	allModels := models.NewModels(ctx, chClient)
	allService := services.NewAllServices(ctx, allModels)

	serverTask := tasks.NewTaskServer(
		allService.Statistic,
		cfg.RedisHost,
		asynq.Config{
			Concurrency: 1,
		},
	)
	tasks.NewClient(cfg.RedisHost)

	app.New(ctx, cfg, serverTask, fb, allService).Start()
}
