package routes

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

//для получения данных о сессиях посетителей.

type SessionHandlers struct {
	app          *fiber.App
	sessionModel *models.SessionModel
}

func NewSessionHandlers(app *fiber.App, sessionModel *models.SessionModel) SessionHandlers {
	return SessionHandlers{
		app:          app,
		sessionModel: sessionModel,
	}
}

func (sh SessionHandlers) AddHandlers() {
	sh.app.Post("/session/filter", sh.Filter)
	sh.app.Delete("/session/delete/:id/", sh.DeleteById)
	sh.app.Delete("/session/delete/list/:list", sh.DeleteByList)
}

func (sh SessionHandlers) Filter(ctx *fiber.Ctx) error {
	var filter filters.Filter
	body := ctx.Body()
	err := jsoniter.Unmarshal(body, &filter)
	if err != nil {
		ctx.Status(502)
		return err
	}
	err, result := sh.sessionModel.Find(filter)
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

func (sh SessionHandlers) DeleteById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id", -1)
	if err != nil {
		return err
	}
	sh.sessionModel.DeleteById(id)
	return nil
}

func (sh SessionHandlers) DeleteByList(ctx *fiber.Ctx) error {
	return nil
}
