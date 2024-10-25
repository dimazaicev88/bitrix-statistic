package routes

import (
	"bitrix-statistic/internal/entityjson"
	"bitrix-statistic/internal/services"
	"context"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type OptionHandlers struct {
	app        *fiber.App
	allService *services.AllServices
	ctx        context.Context
}

func NewOption(ctx context.Context, app *fiber.App, allService *services.AllServices) OptionHandlers {
	return OptionHandlers{
		app:        app,
		allService: allService,
		ctx:        ctx,
	}
}

func (oh OptionHandlers) AddHandlers() {
	oh.app.Post("/api/v1/options/set", oh.set)
	oh.app.Get("/api/v1/options/get", oh.get)
	oh.app.Delete("/api/v1/option/:uuid/", oh.deleteByUuid)
}

func (oh OptionHandlers) deleteByUuid(ctx *fiber.Ctx) error {

	return nil
}

func (oh OptionHandlers) set(ctx *fiber.Ctx) error {
	var serverOptions entityjson.Options
	err := json.Unmarshal(ctx.Body(), &serverOptions)
	if err != nil {
		logrus.Println(err)
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	if serverOptions.AdvCompany != nil {
		err = oh.allService.Option.SetAdvCompany(serverOptions.AdvCompany)
		if err != nil {
			return ctx.JSON(map[string]any{
				"error": err.Error(),
			})
		}
	}

	return ctx.JSON(map[string]any{
		"error": "",
	})
}

func (oh OptionHandlers) list(ctx *fiber.Ctx) error {
	return nil
}

func (oh OptionHandlers) get(ctx *fiber.Ctx) error {
	return ctx.JSON(oh.allService.Option.GetOptions())
}
