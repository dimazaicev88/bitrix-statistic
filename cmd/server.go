package main

import (
	"bitrix-statistic/internal/app"
	"bitrix-statistic/internal/config"
	"context"
	_ "github.com/go-sql-driver/mysql"
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

	app.New(ctx, config.GetServerConfig()).Start()
}
