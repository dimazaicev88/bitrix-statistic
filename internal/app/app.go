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

	allModels := models.NewModels(app.ctx, chClient)

	searcherService := services.NewSearcher(app.ctx, allModels)
	hitService := services.NewHit(app.ctx, allModels)
	advService := services.NewAdv(app.ctx, allModels, hitService)
	guestService := services.NewGuest(app.ctx, allModels, hitService, advService)
	routes.NewMain(fb).AddHandlers()
	routes.NewAdv(app.ctx, fb, services.NewAdv(app.ctx, allModels, hitService)).AddHandlers()
	routes.NewCountry(app.ctx, fb, services.NewCountry(app.ctx, allModels)).AddHandlers()
	routes.NewGuest(app.ctx, fb, services.NewGuest(app.ctx, allModels, hitService, advService)).AddHandlers()
	routes.NewHit(app.ctx, fb, services.NewHit(app.ctx, allModels)).AddHandlers()
	routes.NewPage(app.ctx, fb, services.NewPage(app.ctx, allModels)).AddHandlers()
	routes.NewPath(app.ctx, fb, services.NewPath(app.ctx, allModels)).AddHandlers()
	routes.NewReferer(app.ctx, fb, services.NewReferer(app.ctx, allModels)).AddHandlers()
	routes.NewSearcher(app.ctx, fb, searcherService).AddHandlers()
	routes.NewSearcherHit(app.ctx, fb, searcherService).AddHandlers()
	routes.NewSession(app.ctx, fb, services.NewSession(app.ctx, allModels)).AddHandlers()
	routes.NewStatEvent(app.ctx, fb, services.NewEvent(app.ctx, allModels)).AddHandlers()
	routes.NewStatistic(app.ctx, fb, services.NewStatistic(app.ctx, allModels)).AddHandlers()
	routes.NewStopList(app.ctx, fb, services.NewStopList(app.ctx, allModels)).AddHandlers()
	routes.NewTraffic(app.ctx, fb, services.NewTraffic(app.ctx, allModels)).AddHandlers()
	routes.NewUserOnline(app.ctx, fb, services.NewUserOnline(app.ctx, allModels)).AddHandlers()

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
