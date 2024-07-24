package routes

import (
	"bitrix-statistic/internal/services"
	"context"
	"github.com/gofiber/fiber/v2"
)

// Page Получения данных о посещенных страницах сайта.
type Page struct {
	fbApp       *fiber.App
	ctx         context.Context
	pageService *services.PageService
}

func NewPage(app *fiber.App, pageService *services.PageService, ctx context.Context) *Page {
	return &Page{
		fbApp:       app,
		ctx:         ctx,
		pageService: pageService,
	}
}

func (oh Page) AddHandlers() {
	oh.fbApp.Post("/v1/page/filter", oh.filter)
}

func (oh Page) filter(ctx *fiber.Ctx) error {
	return nil
}
