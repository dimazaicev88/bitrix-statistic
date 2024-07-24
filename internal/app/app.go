package app

import (
	"bitrix-statistic/internal/cache"
	"bitrix-statistic/internal/config"
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
	ctx context.Context
	cfg config.ServerEnvConfig
}

func New(ctx context.Context, cfg config.ServerEnvConfig) *App {
	return &App{
		ctx: ctx,
		cfg: cfg,
	}
}

func (app *App) Start() {
	errStartServer := make(chan error)

	fb := fiber.New()
	chClient, err := storage.NewClickHouseClient(config.GetServerConfig())

	if err != nil {
		logrus.Fatal(err)
	}

	searcherService := services.NewSearcher(app.ctx, chClient)

	routes.NewAdv(app.ctx, fb, services.NewAdv(app.ctx, chClient)).AddHandlers()
	routes.NewCountry(app.ctx, fb, services.NewCountry(app.ctx, chClient)).AddHandlers()
	routes.NewGuest(app.ctx, fb, services.NewGuest(app.ctx, chClient)).AddHandlers()
	routes.NewHit(app.ctx, fb, services.NewHit(app.ctx, chClient)).AddHandlers()
	routes.NewPage(app.ctx, fb, services.NewPage(app.ctx, chClient)).AddHandlers()
	routes.NewPath(app.ctx, fb, services.NewPath(app.ctx, chClient)).AddHandlers()
	routes.NewReferer(app.ctx, fb, services.NewReferer(app.ctx, chClient)).AddHandlers()
	routes.NewSearcher(app.ctx, fb, searcherService).AddHandlers()
	routes.NewSearcherHit(app.ctx, fb, searcherService).AddHandlers()
	routes.NewSession(app.ctx, fb, services.NewSession(app.ctx, chClient)).AddHandlers()
	routes.NewStatEvent(app.ctx, fb, services.NewEvent(app.ctx, chClient)).AddHandlers()
	routes.NewStatistic(app.ctx, fb, services.NewStatistic(app.ctx, chClient)).AddHandlers()
	routes.NewStopList(app.ctx, fb, services.NewStopList(app.ctx, chClient)).AddHandlers()
	routes.NewTraffic(app.ctx, fb, services.NewTraffic(app.ctx, chClient)).AddHandlers()
	routes.NewUserOnline(app.ctx, fb, services.NewUserOnline(app.ctx, chClient)).AddHandlers()

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
		serverQueue.Shutdown()
		cache.Close()
		return
	case err := <-errStartServer:
		log.Fatal(err)
		return
	}
}
