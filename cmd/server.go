package main

import (
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "root", "24zda#1312", "localhost", 3306, "test")
	//db, err := sqlx.Connect("mysql", dataSource)
	//if err != nil {
	//	log.Panic(err)
	//}
	//
	//app := fiber.New()
	//app.Get("/", func(c *fiber.Ctx) error {
	//	row := db.QueryRow("select id from speed")
	//	var id int
	//	err = row.Scan(&id)
	//	return c.SendString("Hello, World 👋!")
	//})
	//log.Fatal(app.Listen(":3000"))

	//sb := sqlbuilder.Select("id", "name").From("demo.user")
	//sb.Where("a=12").Where("b=12")
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
