package app

import (
	"bitrix-statistic/internal/routes"
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

type Application struct {
	ctx context.Context
}

func New() {
	errStartServer := make(chan error)

	fb := fiber.New()
	routes.New(fb, s.RedisAddress, inspector).RegRoutes()

	//start fiber
	go func() {
		log.Println("starting fiber server on port:", s.serverPort)
		errStartServer <- fb.Listen(":" + strconv.Itoa(s.serverPort))
	}()

	select {
	case err := <-errStartServer:
		log.Fatal(err)
		return
	}
}
