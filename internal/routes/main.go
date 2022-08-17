package routes

import (
	"github.com/gofiber/fiber/v2"
)

type MainPageHandler struct {
	app *fiber.App
}

func NewMainPageHandlers(app *fiber.App) MainPageHandler {
	return MainPageHandler{
		app: app,
	}
}

func (mph MainPageHandler) index(c *fiber.Ctx) error {
	return c.Render("index", nil)
}

func (mph MainPageHandler) AddHandler() {
	mph.app.Get("/", mph.index)
}
