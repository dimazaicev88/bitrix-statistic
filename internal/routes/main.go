package routes

import (
	"github.com/gofiber/fiber/v2"
)

type MainPage struct {
	fbApp *fiber.App
}

func NewMain(fbApp *fiber.App) MainPage {
	return MainPage{
		fbApp: fbApp,
	}
}

func (mph MainPage) index(c *fiber.Ctx) error {
	return c.SendString("Statistics server running")
}

func (mph MainPage) AddHandlers() {
	mph.fbApp.Get("/", mph.index)
}
