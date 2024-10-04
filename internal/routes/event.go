package routes

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

// Event Для работы с событиями.
type Event struct {
	fbApp *fiber.App
	ctx   context.Context
}

func NewEvent(ctx context.Context, app *fiber.App) *Event {
	return &Event{
		ctx:   context.Background(),
		fbApp: app,
	}
}

// TODO добавить методы AddByEvents AddCurrent DecodeGID GetGID GetListByGuest
func (e Event) AddHandlers() {
	e.fbApp.Post("/api/v1/event/filter", e.Filter)
	e.fbApp.Delete("/api/v1/event/:uuid/", e.DeleteById)
	e.fbApp.Post("/api/v1/event/set/", e.Add)
}

func (e Event) Filter(ctx *fiber.Ctx) error {
	return nil
}

func (e Event) DeleteById(ctx *fiber.Ctx) error {
	return nil
}

func (e Event) Add(ctx *fiber.Ctx) error {
	return nil
}
