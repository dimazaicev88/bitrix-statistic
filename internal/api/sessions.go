package api

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

type SessionHandlers struct {
	app          *fiber.App
	sessionModel models.SessionModel
}

func NewSessionHandlers(app *fiber.App, sessionModel models.SessionModel) SessionHandlers {
	return SessionHandlers{
		app:          app,
		sessionModel: sessionModel,
	}
}

func (ch SessionHandlers) AddHandlers() {
	ch.app.Post("/city/sessions", ch.Filter)
	ch.app.Post("/city/add", ch.Filter)
	ch.app.Delete("/city/delete/:id/", ch.DeleteById)
}

func (ch SessionHandlers) Filter(ctx *fiber.Ctx) error {
	var filter filters.Filter
	body := ctx.Body()
	err := jsoniter.Unmarshal(body, &filter)
	if err != nil {
		ctx.Status(502)
		return err
	}
	err, result := ch.sessionModel.Find(filter)
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

func (ch SessionHandlers) DeleteById(ctx *fiber.Ctx) error {

	return nil
}
