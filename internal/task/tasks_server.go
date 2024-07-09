package tasks

import (
	"bitrix-statistic/internal/worker"
	"github.com/hibiken/asynq"
	"log"
)

func NewTaskServer(redisAddr string, cfg asynq.Config) (*asynq.Server, *asynq.ServeMux) {
	log.Println("init task server.")
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		cfg,
	)
	mux := asynq.NewServeMux()
	mux.HandleFunc(TaskName, worker.HandleTask)
	return srv, mux
}
