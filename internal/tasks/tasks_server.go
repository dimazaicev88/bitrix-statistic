package tasks

import (
	"context"
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
	//fmt.Println(string(t.Payload()))

	return nil
}
