package tasks

import (
	"bitrix-statistic/internal/worker"
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
	mux.HandleFunc(TaskStatisticAdd, worker.HandleTask)
	return srv, mux
}
