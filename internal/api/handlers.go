package api

import "github.com/gofiber/fiber/v2"

type HitHandlers struct {
	app *fiber.App
}

func NewHitHandlers() HitHandlers {
	return HitHandlers{}
}

func (hh HitHandlers) AddHandlers() {
	hh.app.Post("/filter", hh.Filter)
	hh.app.Post("/delete", hh.DeleteById)
}

func (hh HitHandlers) Filter(ctx *fiber.Ctx) error {
	return nil
}

func (hh HitHandlers) DeleteById(ctx *fiber.Ctx) error {

	return nil
}
