package routes

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

// Phrase Для получения данных по поисковым фразам.
type Phrase struct {
	fbApp *fiber.App
	ctx   context.Context
}

func NewPhrase(ctx context.Context, app *fiber.App) *Phrase {
	return &Phrase{
		ctx:   ctx,
		fbApp: app,
	}
}

func (p Phrase) AddHandlers() {
	p.fbApp.Post("/v1/phrase/filter", p.filter)
}

func (p Phrase) filter(ctx *fiber.Ctx) error {
	return nil
}
