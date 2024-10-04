package routes

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

// EventType Для работы с типами событий.
type EventType struct {
	fbApp *fiber.App
	ctx   context.Context
}

func NewEventType(fbApp *fiber.App, ctx context.Context) *EventType {
	return &EventType{
		fbApp: fbApp,
		ctx:   ctx,
	}
}

// TODO добавит остальные методы
func (et EventType) AddHandlers() {
	et.fbApp.Post("/api/v1/eventType/filter", et.Filter)
	et.fbApp.Delete("/api/v1/eventType/:uuid/", et.DeleteByUuid)
	et.fbApp.Post("/api/v1/eventType/set/", et.Add)
}

func (et EventType) Filter(ctx *fiber.Ctx) error {
	return nil
}

func (et EventType) DeleteByUuid(ctx *fiber.Ctx) error {
	return nil
}

func (et EventType) Add(ctx *fiber.Ctx) error {
	return nil
}
