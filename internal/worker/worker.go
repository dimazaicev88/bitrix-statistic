package worker

import (
	"context"
	"github.com/hibiken/asynq"
)

func HandleTask(ctx context.Context, t *asynq.Task) error {
	//var command entity.Command
	//if err := json.Unmarshal(t.Payload(), &command); err != nil {
	//	return err
	//}
	//cmd := exec.CommandContext(ctx, command.Name, command.Args...)
	//cmd.Stdout = &ws.StdoutWsWriter{}
	//cmd.Stderr = &ws.WriterWsError{}
	//
	//if err := cmd.Start(); err != nil {
	//	return err
	//}
	//
	//if err := cmd.Wait(); err != nil {
	//	return err
	//}

	return nil
}
