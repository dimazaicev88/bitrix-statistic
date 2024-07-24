package app

import (
	"bitrix-statistic/internal/cache"
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

	routes.NewAdv(services.NewAdv(app.ctx, chClient), fb, app.ctx).AddHandlers()
	routes.NewCountry(fb, services.NewCountry(app.ctx, chClient)).AddHandlers()
	routes.NewGuest(fb, services.NewGuest(app.ctx, chClient)).AddHandlers()
	routes.NewHit(fb, services.NewHit(app.ctx, chClient)).AddHandlers()
	routes.NewPage(fb, services.NewPageService(app.ctx, chClient), app.ctx).AddHandlers()

	routes.NewStatistic(fb).RegRoutes()

	routes.NewSessionHandlers(fb, models.NewSession(app.ctx, chClient)).AddHandlers()

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
