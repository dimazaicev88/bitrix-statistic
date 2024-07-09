package app

import (
	"bitrix-statistic/internal/routes"
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

type App struct {
	ctx          context.Context
	RedisAddress string
	serverPort   int
}

func New(redisAddress string, serverPort int) *App {
	return &App{
		ctx:          context.Background(),
		RedisAddress: redisAddress,
		serverPort:   serverPort,
	}
}

func (app *App) Start() {
	errStartServer := make(chan error)

	fb := fiber.New()

	routes.NewStatistic(fb).RegRoutes()

	//start fiber
	go func() {
		log.Println("starting fiber server on port:", app.serverPort)
		errStartServer <- fb.Listen(":" + strconv.Itoa(app.serverPort))
	}()

	select {
	case err := <-errStartServer:
		log.Fatal(err)
		return
	}
}
