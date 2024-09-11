package routes

import (
	"bitrix-statistic/internal/services"
	"context"
	"github.com/gofiber/fiber/v2"
)

// Page Получения данных о посещенных страницах сайта.
type Page struct {
	fbApp      *fiber.App
	ctx        context.Context
	allService *services.AllServices
}

func NewPage(ctx context.Context, app *fiber.App, allService *services.AllServices) *Page {
	return &Page{
		fbApp:      app,
		ctx:        ctx,
		allService: allService,
	}
}

func (oh Page) AddHandlers() {
	oh.fbApp.Post("/v1/page/filter", oh.filter)
}

func (oh Page) filter(ctx *fiber.Ctx) error {
	return nil
}
