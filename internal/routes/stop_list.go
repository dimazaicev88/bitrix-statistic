package routes

import (
	"bitrix-statistic/internal/services"
	"context"
	"github.com/gofiber/fiber/v2"
)

// StopList Для работы со стоп-листом.
type StopList struct {
	fbApp      *fiber.App
	ctx        context.Context
	allService *services.AllServices
}

func NewStopList(ctx context.Context, app *fiber.App, allService *services.AllServices) *StopList {
	return &StopList{
		ctx:        ctx,
		fbApp:      app,
		allService: allService,
	}
}

func (sl StopList) AddHandlers() {
	sl.fbApp.Post("/api/v1/stopList/filter", sl.filter)
	sl.fbApp.Get("/api/v1/stopList/:uuid", sl.findByUuid)
}

func (sl StopList) filter(ctx *fiber.Ctx) error {
	return nil
}

func (sl StopList) findByUuid(ctx *fiber.Ctx) error {
	return nil
}
