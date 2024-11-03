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

// Add TODO добавить отправку json с текстом ошибки.
func (sh *Statistic) Add(ctx *fiber.Ctx) error {
	var userData dto.UserData
	err := json.Unmarshal(ctx.Body(), &userData)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	resultJson, _ := json.Marshal(userData)

	task := asynq.NewTask(tasks.TaskStatisticAdd, resultJson, asynq.MaxRetry(0))
	_, err = tasks.GetClient().EnqueueContext(sh.ctx, task)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(dto.Response{
		Result: "",
		Error:  "",
	})
}
