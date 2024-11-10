package app

import (
	"bitrix-statistic/internal/config"
	"bitrix-statistic/internal/routes"
	"bitrix-statistic/internal/tasks"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/mux"
	"github.com/hibiken/asynq"
	"github.com/hibiken/asynqmon"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"strconv"
)

type App struct {
	fb         *fiber.App
	taskServer *tasks.TaskServer
	cfg        config.ServerEnvConfig
}

func New(
	cfg config.ServerEnvConfig,
	taskServer *tasks.TaskServer,
	fb *fiber.App,
) *App {
	return &App{
		fb:         fb,
		cfg:        cfg,
		taskServer: taskServer,
	}
}

func (app *App) Start(ctx context.Context) {
	errStartServer := make(chan error)

	routes.NewMain(app.fb).AddHandlers()
	routes.NewStatistic(app.fb, ctx).AddHandlers()

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

	h := asynqmon.New(asynqmon.Options{
		RootPath:     "/monitoring",
		RedisConnOpt: asynq.RedisClientOpt{Addr: app.cfg.RedisHost},
	})

	r := mux.NewRouter()
	r.PathPrefix(h.RootPath()).Handler(h)

	srv := &http.Server{
		Handler: r,
		Addr:    ":" + strconv.Itoa(app.cfg.TaskMonitorPort),
	}

	go func() {
		logrus.Println("starting monitoring  server on port:", app.cfg.TaskMonitorPort)
		errStartServer <- srv.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		app.taskServer.AsynqServer.Shutdown()
		tasks.Close()
		return
	case err := <-errStartServer:
		log.Fatal(err)
		return
	}
}
