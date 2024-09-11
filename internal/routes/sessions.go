package routes

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/services"
	"context"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

// SessionHandlers Для получения данных о сессиях посетителей.
type SessionHandlers struct {
	app        *fiber.App
	allService *services.AllServices
	ctx        context.Context
}

func NewSession(ctx context.Context, app *fiber.App, allService *services.AllServices) SessionHandlers {
	return SessionHandlers{
		app:        app,
		allService: allService,
		ctx:        ctx,
	}
}

func (sh SessionHandlers) AddHandlers() {
	sh.app.Post("/session/filter", sh.Filter)
	sh.app.Get("/session/:uuid", sh.findByUuid)
}

func (sh SessionHandlers) Filter(ctx *fiber.Ctx) error {
	var filter filters.Filter
	body := ctx.Body()
	err := json.Unmarshal(body, &filter)
	if err != nil {
		ctx.Status(502)
		return err
	}
	result, err := sh.allService.Session.Filter(filter)
	if err != nil {
		ctx.Status(502)
		return err
	}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		ctx.Status(502)
		return err
	}
	return ctx.SendString(string(jsonResult))
}

func (sh SessionHandlers) DeleteByList(ctx *fiber.Ctx) error {
	return nil
}

func (sh SessionHandlers) findByUuid(ctx *fiber.Ctx) error {
	return nil
}
