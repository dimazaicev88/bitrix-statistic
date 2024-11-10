package main

import (
	"bitrix-statistic/internal/app"
	"bitrix-statistic/internal/config"
	"bitrix-statistic/internal/repository"
	"bitrix-statistic/internal/services"
	"bitrix-statistic/internal/storage"
	"bitrix-statistic/internal/tasks"
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/hibiken/asynq"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file", err.Error())
	}
	cfg := config.GetServerConfig()

	fb := fiber.New()
	fb.Use(cors.New(cors.Config{
		AllowOrigins: "*",                                           // Allow all origins
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",      // Allowed methods
		AllowHeaders: "Origin, Content-Type, Accept, Authorization", // Allowed headers
	}))

	chClient := storage.NewClickHouseClient(config.GetServerConfig())
	if err != nil {
		logrus.Fatal(err)
	}

	guestModel := repository.NewGuest(chClient)
	hitModel := repository.NewHit(chClient)
	serverTask := tasks.NewTaskServer(
		services.NewStatistic(services.NewGuest(guestModel), services.NewHit(hitModel)),
		cfg.RedisHost,
		asynq.Config{
			Concurrency: 1,
		},
	)
	tasks.NewClient(cfg.RedisHost)
	app.New(cfg, serverTask, fb).Start(ctx)
}
