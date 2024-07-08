package app

import (
	"bitrix-statistic/internal/routes"
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

type App struct {
	ctx context.Context
}

func New() {
	errStartServer := make(chan error)

	fb := fiber.New()

	routes.NewStatistic(fb).RegRoutes()

	//start fiber
	go func() {
		log.Println("starting fiber server on port:", 9007)
		errStartServer <- fb.Listen(":" + strconv.Itoa(9007))
	}()

	select {
	case err := <-errStartServer:
		log.Fatal(err)
		return
	}
}
