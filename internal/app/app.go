package app

import (
	"bitrix-statistic/internal/config"
	"bitrix-statistic/internal/routes"
	"bitrix-statistic/internal/tasks"
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

type App struct {
	ctx        context.Context
	fb         *fiber.App
	taskServer *tasks.TaskServer
	cfg        config.ServerEnvConfig
}

func New(
	ctx context.Context,
	cfg config.ServerEnvConfig,
	taskServer *tasks.TaskServer,
	fb *fiber.App,
) *App {
	return &App{
		ctx:        ctx,
		fb:         fb,
		cfg:        cfg,
		taskServer: taskServer,
	}
}

func (app *App) Start() {
	errStartServer := make(chan error)

	routes.NewMain(app.fb).AddHandlers()
	routes.NewStatistic(app.ctx, app.fb, app.AllServices).AddHandlers()
	routes.NewStopList(app.ctx, app.fb, app.AllServices).AddHandlers()

	//start fiber
	go func() {
		log.Println("starting fiber server on port:", app.cfg.ServerPort)
		errStartServer <- app.fb.Listen(":" + strconv.Itoa(app.cfg.ServerPort))
	}()

	//start server queue
	go func() {
		log.Println("starting task server")
		errStartServer <- app.taskServer.AsynqServer.Run(app.taskServer.AsynqServeMux)
	}()

	select {
	case <-app.ctx.Done():
		app.taskServer.AsynqServer.Shutdown()
		return
	case err := <-errStartServer:
		log.Fatal(err)
		return
	}
}
