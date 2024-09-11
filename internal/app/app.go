package app

import (
	"bitrix-statistic/internal/config"
	"bitrix-statistic/internal/routes"
	"bitrix-statistic/internal/services"
	"bitrix-statistic/internal/tasks"
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

type App struct {
	ctx         context.Context
	AllServices *services.AllServices
	fb          *fiber.App
	taskServer  *tasks.TaskServer
	cfg         config.ServerEnvConfig
}

func New(
	ctx context.Context,
	cfg config.ServerEnvConfig,
	taskServer *tasks.TaskServer,
	fb *fiber.App,
	allServices *services.AllServices,
) *App {
	return &App{
		ctx:         ctx,
		fb:          fb,
		cfg:         cfg,
		taskServer:  taskServer,
		AllServices: allServices,
	}
}

func (app *App) Start() {
	errStartServer := make(chan error)

	routes.NewMain(app.fb).AddHandlers()
	routes.NewAdv(app.ctx, app.fb, app.AllServices).AddHandlers()
	routes.NewCountry(app.ctx, app.fb, app.AllServices).AddHandlers()
	routes.NewGuest(app.ctx, app.fb, app.AllServices).AddHandlers()
	routes.NewHit(app.ctx, app.fb, app.AllServices).AddHandlers()
	routes.NewPage(app.ctx, app.fb, app.AllServices).AddHandlers()
	routes.NewPath(app.ctx, app.fb, app.AllServices).AddHandlers()
	routes.NewReferer(app.ctx, app.fb, app.AllServices).AddHandlers()
	routes.NewSearcher(app.ctx, app.fb, app.AllServices).AddHandlers()
	routes.NewSearcherHit(app.ctx, app.fb, app.AllServices).AddHandlers()
	routes.NewSession(app.ctx, app.fb, app.AllServices).AddHandlers()
	routes.NewStatEvent(app.ctx, app.fb, app.AllServices).AddHandlers()
	routes.NewStatistic(app.ctx, app.fb, app.AllServices).AddHandlers()
	routes.NewStopList(app.ctx, app.fb, app.AllServices).AddHandlers()
	routes.NewTraffic(app.ctx, app.fb, app.AllServices).AddHandlers()
	routes.NewUserOnline(app.ctx, app.fb, app.AllServices).AddHandlers()

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
		app.AllServices.Guest.ClearCache()
		app.AllServices.Statistic.ClearCache()
		app.taskServer.AsynqServer.Shutdown()
		return
	case err := <-errStartServer:
		log.Fatal(err)
		return
	}
}
