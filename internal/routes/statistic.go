package routes

import (
	"bitrix-statistic/internal/dto"
	"bitrix-statistic/internal/tasks"
	"context"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/hibiken/asynq"
)

type Statistic struct {
	ctx   context.Context
	fbApp *fiber.App
}

type Answer struct {
	Msg string `json:"msg"`
}

func NewStatistic(fbApp *fiber.App, ctx context.Context) *Statistic {
	return &Statistic{
		ctx:   ctx,
		fbApp: fbApp,
	}
}

func (sh *Statistic) AddHandlers() {
	sh.fbApp.Post("/api/v1/statistic/add", sh.Add)
}

// Add Добавить задачу в очередь
func (sh *Statistic) Add(ctx *fiber.Ctx) error {
	var userData dto.UserData
	if err := json.Unmarshal(ctx.Body(), &userData); err != nil {
		return ctx.JSON(dto.Response{
			Error: err.Error(),
		})
	}

	task := asynq.NewTask(tasks.TaskStatisticAdd, ctx.Body(), asynq.MaxRetry(0))
	if _, err := tasks.GetClient().EnqueueContext(sh.ctx, task); err != nil {
		return ctx.JSON(dto.Response{
			Error: err.Error(),
		})
	}

	return nil
}
