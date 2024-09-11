package app

import (
	"bitrix-statistic/internal/config"
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/routes"
	"bitrix-statistic/internal/services"
	"bitrix-statistic/internal/storage"
	"bitrix-statistic/internal/tasks"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
	"log"
	"strconv"
)

type App struct {
	ctx         context.Context
	cfg         config.ServerEnvConfig
	AllServices *services.AllServices
}

func New(ctx context.Context) *App {
	return &App{
		ctx: ctx,
	}
}

func (app *App) Start() {
	errStartServer := make(chan error)

	fb := fiber.New()
	chClient, err := storage.NewClickHouseClient(config.GetServerConfig())

	if err != nil {
		logrus.Fatal(err)
	}

	allModels := models.NewModels(app.ctx, chClient)

	allService := services.NewAllServices(app.ctx, allModels)
	routes.NewMain(fb).AddHandlers()
	routes.NewAdv(app.ctx, fb, allService).AddHandlers()
	routes.NewCountry(app.ctx, fb, allService).AddHandlers()
	routes.NewGuest(app.ctx, fb, allService).AddHandlers()
	routes.NewHit(app.ctx, fb, allService).AddHandlers()
	routes.NewPage(app.ctx, fb, allService).AddHandlers()
	routes.NewPath(app.ctx, fb, allService).AddHandlers()
	routes.NewReferer(app.ctx, fb, allService).AddHandlers()
	routes.NewSearcher(app.ctx, fb, allService).AddHandlers()
	routes.NewSearcherHit(app.ctx, fb, allService).AddHandlers()
	routes.NewSession(app.ctx, fb, allService).AddHandlers()
	routes.NewStatEvent(app.ctx, fb, allService).AddHandlers()
	routes.NewStatistic(app.ctx, fb, allService).AddHandlers()
	routes.NewStopList(app.ctx, fb, allService).AddHandlers()
	routes.NewTraffic(app.ctx, fb, allService).AddHandlers()
	routes.NewUserOnline(app.ctx, fb, allService).AddHandlers()

	//start fiber
	go func() {
		log.Println("starting fiber server on port:", app.cfg.ServerPort)
		errStartServer <- fb.Listen(":" + strconv.Itoa(app.cfg.ServerPort))
	}()

	tasks.NewClient(app.cfg.RedisHost)
	serverQueue, serverMux := tasks.NewTaskServer(
		app.cfg.RedisHost,
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
		allService.Guest.ClearCache()
		allService.Statistic.ClearCache()
		serverQueue.Shutdown()
		return
	case err := <-errStartServer:
		log.Fatal(err)
		return
	}
}
