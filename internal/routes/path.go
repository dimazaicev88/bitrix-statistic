package routes

import (
	"bitrix-statistic/internal/services"
	"context"
	"github.com/gofiber/fiber/v2"
)

// Path Для получения данных о путях по сайту.
type Path struct {
	fbApp       *fiber.App
	ctx         context.Context
	pathService *services.PathService
}

func NewPath(ctx context.Context, app *fiber.App, pathService *services.PathService) *Path {
	return &Path{
		fbApp:       app,
		ctx:         ctx,
		pathService: pathService,
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
