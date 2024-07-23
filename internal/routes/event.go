package routes

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

//для работы с событиями.

type Event struct {
	fbApp *fiber.App
	ctx   context.Context
}

func NewEvent(ctx context.Context, app *fiber.App) *Event {
	return &Event{
		ctx:   context.Background(),
		fbApp: app,
	}
}
