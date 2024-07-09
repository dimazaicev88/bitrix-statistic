package app

import (
	"bitrix-statistic/internal/routes"
	"bitrix-statistic/internal/tasks"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/hibiken/asynq"
	"log"
	"strconv"
)

type App struct {
	ctx          context.Context
	redisAddress string
	serverPort   int
}

func New(ctx context.Context, redisAddress string, serverPort int) *App {
	return &App{
		ctx:          ctx,
		redisAddress: redisAddress,
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

	tasks.NewClient(app.redisAddress)
	serverQueue, serverMux := tasks.NewTaskServer(
		app.redisAddress,
		asynq.Config{
			Concurrency: 1,
		},
	)

	//start server queue
	go func() {
		log.Println("starting task server")
		errStartServer <- serverQueue.Run(serverMux)
	}()

	select {
	case <-app.ctx.Done():
		serverQueue.Shutdown()
		return
	case err := <-errStartServer:
		log.Fatal(err)
		return
	}
}
