package app

import (
	"context"
	"log"
)

type Application struct {
	ctx context.Context
}

func New() {
	errStartServer := make(chan error)

	//fb := fiber.New()
	//routes.NewStatistic(fb, s.RedisAddress, inspector).RegRoutes()

	//start fiber
	//go func() {
	//	log.Println("starting fiber server on port:", s.serverPort)
	//	errStartServer <- fb.Listen(":" + strconv.Itoa(s.serverPort))
	//}()

	select {
	case err := <-errStartServer:
		log.Fatal(err)
		return
	}
}
