package tasks

import (
	"bitrix-statistic/internal/dto"
	"bitrix-statistic/internal/services"
	"context"
	"github.com/goccy/go-json"
	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
)

const TaskStatisticAdd = "statistic:add"

type TaskServer struct {
	serviceStatisticS *services.Statistic
	AsynqServer       *asynq.Server
	AsynqServeMux     *asynq.ServeMux
}

func NewTaskServer(serviceStatisticS *services.Statistic, redisAddr string, cfg asynq.Config) *TaskServer {
	logrus.Infoln("init tasks server.")
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		cfg,
	)
	mux := asynq.NewServeMux()
	ts := &TaskServer{
		serviceStatisticS: serviceStatisticS,
		AsynqServer:       srv,
		AsynqServeMux:     mux,
	}
	mux.HandleFunc(TaskStatisticAdd, ts.HandleTask)
	return ts
}

func (ts TaskServer) HandleTask(ctx context.Context, t *asynq.Task) error {
	var userData dto.UserData

	if err := json.Unmarshal(t.Payload(), &userData); err != nil {
		return err
	}

	if err := ts.serviceStatisticS.Add(ctx, userData, false); err != nil {
		return err
	}
	return nil
}
