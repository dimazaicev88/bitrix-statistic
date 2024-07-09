package worker

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
)

func HandleTask(ctx context.Context, t *asynq.Task) error {
	fmt.Println(string(t.Payload()))

	return nil
}
