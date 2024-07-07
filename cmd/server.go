package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	//ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	//defer stop()

	//dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "root", "24zda#1312", "localhost", 3306, "test")
	//db, err := sqlx.Connect("mysql", dataSource)
	//if err != nil {
	//	log.Panic(err)
	//}
	//
	app := fiber.New()
	app.Post("/", func(c *fiber.Ctx) error {

		fmt.Println(string(c.Body()))

		//row := db.QueryRow("select id from speed")
		//var id int
		//err = row.Scan(&id)
		return c.SendString("Hello, World 👋!")
	})
	log.Fatal(app.Listen(":3000"))

	//sb := sqlbuilder.Select("id", "name").From("demo.user")
	//sb.build("a=12").build("b=12")
	//fmt.Println(sb.String())
	//serverConfig := config.ParseServerConfig()
	//mysqlStorage := storage.NewMysqlStorage(serverConfig)
	//serverStatistic := server.NewServer(mysqlStorage)
	//err := serverStatistic.Start(125)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//defer func(serverStatistic *server.Server) {
	//	err := serverStatistic.Stop()
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//}(serverStatistic)
}
