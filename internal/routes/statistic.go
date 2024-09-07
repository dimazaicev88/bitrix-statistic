package routes

import (
	"bitrix-statistic/internal/services"
	"bitrix-statistic/internal/tasks"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/hibiken/asynq"
	"log"
)

type Statistic struct {
	fbApp      *fiber.App
	ctx        context.Context
	allService *services.AllService
}

type Answer struct {
	Msg string `json:"msg"`
}

func NewStatistic(ctx context.Context, fbApp *fiber.App, allService *services.AllService) *Statistic {
	return &Statistic{
		fbApp:      fbApp,
		ctx:        ctx,
		allService: allService,
	}
}

func (sh *Statistic) AddHandlers() {
	sh.fbApp.Post("/statistic/add", sh.Add)
}

func (sh *Statistic) Add(ctx *fiber.Ctx) error {
	//asynq.Timeout(time.Second*8)
	task := asynq.NewTask(tasks.TaskStatisticAdd, ctx.Body(), asynq.MaxRetry(0))
	_, err := tasks.GetClient().EnqueueContext(ctx.Context(), task)
	if err != nil {
		log.Panic(err)
	}
	return ctx.SendStatus(200)
}
