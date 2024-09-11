package routes

import (
	"bitrix-statistic/internal/entityjson"
	"bitrix-statistic/internal/services"
	"bitrix-statistic/internal/tasks"
	"context"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"log"
)

type Statistic struct {
	fbApp      *fiber.App
	ctx        context.Context
	allService *services.AllServices
}

type Answer struct {
	Msg string `json:"msg"`
}

func NewStatistic(ctx context.Context, fbApp *fiber.App, allService *services.AllServices) *Statistic {
	return &Statistic{
		fbApp:      fbApp,
		ctx:        ctx,
		allService: allService,
	}
}

func (sh *Statistic) AddHandlers() {
	sh.fbApp.Post("/statistic/add", sh.Add)
}

// Add TODO добавить отправку json с текстом ошибки.
func (sh *Statistic) Add(ctx *fiber.Ctx) error {
	var userData entityjson.UserData
	err := json.Unmarshal(ctx.Body(), &userData)
	if err != nil {
		log.Println(err)
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	if userData.GuestUuid == uuid.Nil {
		userData.GuestUuid = uuid.New()
	}

	resultJson, _ := json.Marshal(userData)

	task := asynq.NewTask(tasks.TaskStatisticAdd, resultJson, asynq.MaxRetry(0))
	_, err = tasks.GetClient().EnqueueContext(ctx.Context(), task)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	answer, err := json.Marshal(map[string]string{
		"guestUuid": userData.GuestUuid.String(),
	})

	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Send(answer)
}
