package routes

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/services"
	"context"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

// EventType Для работы с типами событий.
type EventType struct {
	fbApp       *fiber.App
	ctx         context.Context
	allServices *services.AllServices
}

func NewEventType(fbApp *fiber.App, ctx context.Context, allServices *services.AllServices) *EventType {
	return &EventType{
		fbApp:       fbApp,
		ctx:         ctx,
		allServices: allServices,
	}
}

// TODO добавит остальные методы
func (et EventType) AddHandlers() {
	et.fbApp.Post("/api/v1/eventType/filter", et.Filter)
	et.fbApp.Delete("/api/v1/eventType/:uuid/", et.DeleteByUuid)
	et.fbApp.Post("/api/v1/eventType/set/", et.Add)
}

func (et EventType) Filter(ctx *fiber.Ctx) error {
	var filter filters.Filter
	body := ctx.Body()
	err := json.Unmarshal(body, &filter)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	result, err := et.allServices.Event.Find(filter) //TODO добавить парсинг
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	return ctx.SendString(string(resultJson))
}

func (et EventType) DeleteByUuid(ctx *fiber.Ctx) error {
	return nil
}

func (et EventType) Add(ctx *fiber.Ctx) error {
	return nil
}
