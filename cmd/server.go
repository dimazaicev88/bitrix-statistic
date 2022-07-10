package main

import (
	"bitrix-statistic/internal/server"
	"log"
)

func main() {
	serverStatistic := server.NewServer()
	err := serverStatistic.Start(3000)
	if err != nil {
		log.Fatalln(err)
	}

	defer func(serverStatistic *server.Server) {
		err := serverStatistic.Stop()
		if err != nil {
			log.Fatalln(err)
		}
	}(serverStatistic)

	//app := fiber.New()
	//app.Get("/", func(c *fiber.Ctx) error {
	//	return c.SendString("Hello, World 👋!")
	//})
	//err = app.Listen(":3000")
	//if err != nil {
	//	log.Fatalln(err)
	//}
}
