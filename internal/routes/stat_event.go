package routes

import (
	"bitrix-statistic/internal/services"
	"context"
	"github.com/gofiber/fiber/v2"
)

// StatEvent Возвращает список хитов поисковых систем.
type StatEvent struct {
	fbApp      *fiber.App
	ctx        context.Context
	allService *services.AllServices
}

func NewStatEvent(ctx context.Context, fbApp *fiber.App, allService *services.AllServices) *StatEvent {
	return &StatEvent{
		fbApp:      fbApp,
		ctx:        ctx,
		allService: allService,
	}
}

func (p StatEvent) AddHandlers() {
	p.fbApp.Post("/api/v1/statEvent/set", p.add)
	p.fbApp.Delete("/api/v1/statEvent/:uuid", p.delete)
}

func (p StatEvent) add(ctx *fiber.Ctx) error {
	return nil
}

func (p StatEvent) delete(ctx *fiber.Ctx) error {
	return nil
}
