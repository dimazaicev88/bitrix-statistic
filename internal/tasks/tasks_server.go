package tasks

import (
	"bitrix-statistic/internal/app"
	"bitrix-statistic/internal/entityjson"
	"context"
	"github.com/goccy/go-json"
	"github.com/hibiken/asynq"
	"log"
)

const TaskStatisticAdd = "statistic:add"
const TaskGroup = "default"

func NewTaskServer(redisAddr string, cfg asynq.Config) (*asynq.Server, *asynq.ServeMux) {
	log.Println("init tasks server.")
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		cfg,
	)
	mux := asynq.NewServeMux()
	mux.HandleFunc(TaskStatisticAdd, HandleTask)
	return srv, mux
}

func HandleTask(ctx context.Context, t *asynq.Task) error {
	var userData entityjson.UserData

	if err := json.Unmarshal(t.Payload(), &userData); err != nil {
		return err
	}

	if err := app.Server().Get().AllServices.Statistic.Add(userData); err != nil {
		return err
	}

	return nil
}
