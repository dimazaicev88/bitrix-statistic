package routes

import (
	"github.com/gofiber/fiber/v2"
)

type MainPageHandler struct {
	fbApp *fiber.App
}

func NewMainPageHandlers(fbApp *fiber.App) MainPageHandler {
	return MainPageHandler{
		fbApp: fbApp,
	}
}

func (mph MainPageHandler) index(c *fiber.Ctx) error {
	return c.Render("index", nil)
}

func (mph MainPageHandler) AddHandler() {
	mph.fbApp.Get("/", mph.index)
}
