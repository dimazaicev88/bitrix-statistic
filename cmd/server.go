package main

import (
	"bitrix-statistic/internal/app"
	"context"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	app.New(ctx, "127.0.0.1:6379", 9008).Start()
}
