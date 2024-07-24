package routes

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/services"
	"context"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

// SessionHandlers Для получения данных о сессиях посетителей.
type SessionHandlers struct {
	app            *fiber.App
	sessionService *services.SessionService
	ctx            context.Context
}

func NewSession(ctx context.Context, app *fiber.App, sessionService *services.SessionService) SessionHandlers {
	return SessionHandlers{
		app:            app,
		sessionService: sessionService,
		ctx:            ctx,
	}
}

func (sh SessionHandlers) AddHandlers() {
	sh.app.Post("/session/filter", sh.Filter)
	sh.app.Get("/session/:uuid", sh.findByUuid)
}

func (sh SessionHandlers) Filter(ctx *fiber.Ctx) error {
	var filter filters.Filter
	body := ctx.Body()
	err := jsoniter.Unmarshal(body, &filter)
	if err != nil {
		ctx.Status(502)
		return err
	}
	result, err := sh.sessionService.Filter(filter)
	if err != nil {
		ctx.Status(502)
		return err
	}

	json, err := jsoniter.MarshalToString(result)
	if err != nil {
		ctx.Status(502)
		return err
	}
	return ctx.SendString(json)
}

func (sh SessionHandlers) DeleteByList(ctx *fiber.Ctx) error {
	return nil
}

func (sh SessionHandlers) findByUuid(ctx *fiber.Ctx) error {
	return nil
}
