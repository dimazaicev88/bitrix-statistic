package routes

import (
	"bitrix-statistic/internal/services"
	"context"
	"github.com/gofiber/fiber/v2"
)

// Path Для получения данных о путях по сайту.
type Path struct {
	fbApp      *fiber.App
	ctx        context.Context
	allService *services.AllService
}

func NewPath(ctx context.Context, app *fiber.App, allService *services.AllService) *Path {
	return &Path{
		fbApp:      app,
		ctx:        ctx,
		allService: allService,
	}
}

func (p Path) AddHandlers() {
	p.fbApp.Post("/v1/path/filter", p.filter)
	p.fbApp.Get("/v1/path/:uuid", p.findBydUuid)
}

func (p Path) filter(ctx *fiber.Ctx) error {
	return nil
}

func (p Path) findBydUuid(ctx *fiber.Ctx) error {
	return nil
}
