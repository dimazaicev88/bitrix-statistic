package routes

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/services"
	"context"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

// Phrase Для получения данных по поисковым фразам.
type Phrase struct {
	fbApp       *fiber.App
	ctx         context.Context
	allServices *services.AllServices
}

func NewPhrase(ctx context.Context, app *fiber.App, allServices *services.AllServices) *Phrase {
	return &Phrase{
		ctx:         ctx,
		fbApp:       app,
		allServices: allServices,
	}
}

func (p Phrase) AddHandlers() {
	p.fbApp.Post("/api/v1/phrase/filter", p.filter)
}

func (p Phrase) filter(ctx *fiber.Ctx) error {
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
