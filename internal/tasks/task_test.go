package tasks

import (
	"context"
	"github.com/hibiken/asynq"
	"log"
	"testing"
)

func BenchmarkAddTask(b *testing.B) {
	redisUrl := "127.0.0.1:6379"
	NewClient(redisUrl)
	serverQueue, serverMux := NewTaskServer(
		redisUrl,
		asynq.Config{
			Concurrency: 1,
		},
	)

	defer serverQueue.Shutdown()

	go func() {
		log.Println("starting task server")
		serverQueue.Run(serverMux)
	}()

	for i := 0; i < b.N; i++ {
		task := asynq.NewTask(TaskStatisticAdd, []byte("test"), asynq.MaxRetry(0))
		GetClient().EnqueueContext(context.Background(), task)
	}
}
