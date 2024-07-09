package routes

import (
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/tasks"
	"github.com/gofiber/fiber/v2"
	"github.com/hibiken/asynq"
	"log"
)

type Statistic struct {
	fbApp          *fiber.App
	statisticModel models.StatisticModel
}

type Answer struct {
	Msg string `json:"msg"`
}

func NewStatistic(fbApp *fiber.App) *Statistic {
	return &Statistic{
		fbApp: fbApp,
	}
}

func (sh *Statistic) RegRoutes() {
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
