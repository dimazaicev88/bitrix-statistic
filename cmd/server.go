package main

import (
	"bitrix-statistic/internal/app"
	"bitrix-statistic/internal/config"
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/services"
	"bitrix-statistic/internal/storage"
	"bitrix-statistic/internal/tasks"
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"github.com/huandu/go-sqlbuilder"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//TODO для  всех методов добавить по возможности одинаковые ошибки.

func main() {
	sqlbuilder.DefaultFlavor = sqlbuilder.ClickHouse
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file", err.Error())
	}
	cfg := config.GetServerConfig()

	fb := fiber.New()
	fb.Use(cors.New(cors.Config{
		AllowOrigins: "*",                                           // Allow all origins
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",      // Allowed methods
		AllowHeaders: "Origin, Content-Type, Accept, Authorization", // Allowed headers
	}))

	chClient, err := storage.NewClickHouseClient(config.GetServerConfig())
	if err != nil {
		logrus.Fatal(err)
	}
	// Prepare for batch insert
	batch, err := chClient.PrepareBatch(context.Background(), `INSERT INTO hit (uuid, session_uuid, adv_uuid, date_hit, php_session_id, guest_uuid, language, is_new_guest, user_id, user_auth, url,
		url_404, url_from, ip, method, cookies, user_agent, stop_list_uuid, country_id, city_uuid, site_id, favorites) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)

	if err != nil {
		log.Fatalf("Failed to prepare batch insert: %v", err)
	}

	hitUuidVal := uuid.New()
	sessionUuidVal := uuid.New()
	advUuidVal := uuid.New()
	guestUuidVal := uuid.New()

	// Insert data in batches
	for i := 0; i < 5000000; i++ {
		if err := batch.Append(
			hitUuidVal,                 // uuid
			sessionUuidVal,             // session_uuid
			advUuidVal,                 // adv_uuid
			time.Now(),                 // date_hit
			"sessionID123",             // php_session_id
			guestUuidVal,               // guest_uuid
			"en",                       // language
			false,                      // is_new_guest
			uint32(i),                  // user_id
			false,                      // user_auth
			"http://example.com",       // url
			false,                      // url_404
			"http://referrer.com",      // url_from
			net.ParseIP("192.168.1.1"), // ip
			"GET",                      // method
			"cookie_data",              // cookies
			"Mozilla/5.0",              // user_agent
			uuid.New(),                 // stop_list_uuid
			"RU",                       // country_id (FixedString(2))
			uuid.New(),                 // city_uuid
			"01",                       // site_id (FixedString(2))
			false,                      // favorites
		); err != nil {
			log.Fatalf("Failed to append data to batch: %v", err)
		}

		////Flush after every 1000 rows for better performance.
		//if (i+1)%50000 == 0 {
		//	if err := batch.Send(); err != nil {
		//		log.Fatalf("Failed to send batch: %v", err)
		//	}
		//}
	}
	if err := batch.Send(); err != nil {
		log.Fatalf("Failed to send batch: %v", err)
	}

	// Flush any remaining data.
	if err := batch.Send(); err != nil {
		log.Fatalf("Failed to send final batch: %v", err)
	}
	//batch, err := chClient.PrepareBatch(context.Background(), "INSERT INTO hit")
	//if err != nil {
	//	logrus.Fatal(err)
	//}
	//
	//var (
	//	hitUuid        []uuid.UUID
	//	session_uuid   []uuid.UUID
	//	adv_uuid       []uuid.UUID
	//	date_hit       []time.Time
	//	php_session_id []string
	//	guest_uuid     []uuid.UUID
	//	language       []string
	//	is_new_guest   []bool
	//	user_id        []uint32
	//	user_auth      []bool
	//	url            []string
	//	url_404        []bool
	//	url_from       []string
	//	ip             []string
	//	method         []string
	//	cookies        []string
	//	user_agent     []string
	//	stop_list_uuid []uuid.UUID
	//	country_id     []string
	//	city_uuid      []uuid.UUID
	//	site_id        []string
	//)

	allModels := models.NewModels(ctx, chClient)
	allService := services.NewAllServices(ctx, allModels)

	serverTask := tasks.NewTaskServer(
		allService.Statistic,
		cfg.RedisHost,
		asynq.Config{
			Concurrency: 1,
		},
	)
	tasks.NewClient(cfg.RedisHost)

	app.New(ctx, cfg, serverTask, fb, allService).Start()
}
