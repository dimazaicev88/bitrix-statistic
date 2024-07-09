package tasks

import (
	"github.com/hibiken/asynq"
	"log"
	"sync"
)

var (
	client *asynq.Client
	once   sync.Once
)

func NewClient(redisAddress string) {
	once.Do(func() {
		log.Println("Create redis client. Redis address:", redisAddress)
		client = asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddress})
	})
}

func Close() {
	if client != nil {
		client.Close()
	}
}

func GetClient() *asynq.Client {
	return client
}
