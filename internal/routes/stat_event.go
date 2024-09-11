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
	p.fbApp.Post("/v1/statEvent/filter", p.filter)
	p.fbApp.Post("/v1/statEvent/filterByGuest", p.findByGuest)
	p.fbApp.Post("/v1/statEvent/add", p.add)
	p.fbApp.Post("/v1/statEvent/addByEvents", p.addByEvents)
	p.fbApp.Post("/v1/statEvent/addCurrent", p.addCurrent)
	p.fbApp.Post("/v1/statEvent/decodeGid/:gid", p.decodeGid)
	p.fbApp.Delete("/v1/statEvent/:uuid", p.delete)
	p.fbApp.Get("/v1/statEvent/gid/:gid", p.findGid)
}

func (p StatEvent) filter(ctx *fiber.Ctx) error {
	return nil
}

func (p StatEvent) add(ctx *fiber.Ctx) error {
	return nil
}

func (p StatEvent) addByEvents(ctx *fiber.Ctx) error {
	return nil
}

func (p StatEvent) addCurrent(ctx *fiber.Ctx) error {
	return nil
}

func (p StatEvent) decodeGid(ctx *fiber.Ctx) error {
	return nil
}

func (p StatEvent) delete(ctx *fiber.Ctx) error {
	return nil
}

func (p StatEvent) findGid(ctx *fiber.Ctx) error {
	return nil
}

func (p StatEvent) findByGuest(ctx *fiber.Ctx) error {
	return nil
}
