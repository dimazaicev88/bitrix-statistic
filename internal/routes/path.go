package routes

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/services"
	"context"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

// Path Для получения данных о путях по сайту.
type Path struct {
	fbApp       *fiber.App
	ctx         context.Context
	allServices *services.AllServices
}

func NewPath(ctx context.Context, app *fiber.App, allService *services.AllServices) *Path {
	return &Path{
		fbApp:       app,
		ctx:         ctx,
		allServices: allService,
	}
}

func (p Path) AddHandlers() {
	p.fbApp.Post("/api/v1/path/filter", p.filter)
	p.fbApp.Get("/api/v1/path/:uuid", p.findBydUuid)
}

func (p Path) filter(ctx *fiber.Ctx) error {
	var filter filters.Filter
	body := ctx.Body()
	err := json.Unmarshal(body, &filter)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	result, err := p.allServices.Path.Find(filter)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	return ctx.SendString(string(resultJson))
}

func (p Path) findBydUuid(ctx *fiber.Ctx) error {
	return nil
}
