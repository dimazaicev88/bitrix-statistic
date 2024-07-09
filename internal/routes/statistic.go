package routes

import (
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/tasks"
	"github.com/gofiber/fiber/v2"
	"github.com/hibiken/asynq"
	"time"
)

type Statistic struct {
	fbApp          *fiber.App
	statisticModel models.StatisticModel
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
	asynq.NewTask(tasks.TaskStatisticAdd, ctx.Body(), asynq.MaxRetry(0), asynq.Timeout(time.Hour*8))
	return nil
}
