package routes

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

// Page Получения данных о посещенных страницах сайта.
type Page struct {
	fbApp *fiber.App
	ctx   context.Context
}

func NewPage(ctx context.Context, app *fiber.App) *Page {
	return &Page{
		fbApp: app,
		ctx:   ctx,
	}
}

func (oh Page) AddHandlers() {
	oh.fbApp.Post("/v1/page/filter", oh.filter)
}

func (oh Page) filter(ctx *fiber.Ctx) error {
	return nil
}
