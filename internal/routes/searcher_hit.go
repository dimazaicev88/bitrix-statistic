package routes

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

// SearcherHit Для получения данных о хитах поисковых систем (про индексированных страниц).
type SearcherHit struct {
	fbApp *fiber.App
	ctx   context.Context
}

func NewSearcherHit(fbApp *fiber.App, ctx context.Context) *SearcherHit {
	return &SearcherHit{
		fbApp: fbApp,
		ctx:   ctx,
	}
}

func (p SearcherHit) AddHandlers() {
	p.fbApp.Post("/v1/searcherHit/filter", p.filter)
}

func (p SearcherHit) filter(ctx *fiber.Ctx) error {
	return nil
}
