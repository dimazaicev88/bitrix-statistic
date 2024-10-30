package routes

import (
	"bitrix-statistic/internal/services"
	"context"
	"github.com/gofiber/fiber/v2"
)

// Event Для работы с событиями.
type Event struct {
	fbApp       *fiber.App
	ctx         context.Context
	allServices *services.AllServices
}

func NewEvent(ctx context.Context, app *fiber.App, allServices *services.AllServices) *Event {
	return &Event{
		ctx:         ctx,
		fbApp:       app,
		allServices: allServices,
	}
}

func (e Event) AddHandlers() {
	e.fbApp.Delete("/api/v1/event/:uuid/", e.delete)
	e.fbApp.Post("/api/v1/event/add/", e.Add)
}

func (e Event) delete(ctx *fiber.Ctx) error {
	err := e.allServices.Event.Delete(ctx.Params("uuid"))
	if err != nil {
		return err
	}
	return nil
}

func (e Event) Add(ctx *fiber.Ctx) error {
	return nil
}
