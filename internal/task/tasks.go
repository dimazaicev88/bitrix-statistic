package tasks

import (
	"bitrix-statistic/internal/entity"
	"encoding/json"
	"github.com/hibiken/asynq"
	"time"
)

const TaskName = "statistic:add"

func NewTask(name string, args []string, taskName string) (*asynq.Task, error) {
	payload, err := json.Marshal(entity.StatData{})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(taskName, payload, asynq.MaxRetry(0), asynq.Timeout(time.Hour*8)), nil
}
