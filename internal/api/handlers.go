package api

import "github.com/gofiber/fiber/v2"

type HitHandlers struct {
	app *fiber.App
}

func NewHitHandlers(app *fiber.App) HitHandlers {
	return HitHandlers{
		app: app,
	}
}

func (hh HitHandlers) AddHandlers() {
	hh.app.Post("/hit/filter", hh.Filter)
	hh.app.Post("/hit/delete", hh.DeleteById)
}

func (hh HitHandlers) Filter(ctx *fiber.Ctx) error {
	var mapFilter map[string]interface{}
	err := ctx.BodyParser(&mapFilter)
	if err != nil {
		ctx.Status(502)
		return err
	}
	return ctx.SendString("Hit filter")
}

func (hh HitHandlers) DeleteById(ctx *fiber.Ctx) error {

	return nil
}
