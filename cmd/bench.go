package main

import (
	"context"
	"fmt"
	"github.com/centrifugal/gocent/v3"
	"log"
)

func main() {
	c := gocent.New(gocent.Config{
		Addr: "https://centrifugo.intsite.org/api",
		Key:  "",
	})

	ch := "cicd/portal/chat"
	ctx := context.Background()
	//start := time.Now()
	publish, err := c.Publish(ctx, ch, []byte(`{"user": "zabbix","msg": "test message","chatId": 1924358}`))
	if err != nil {
		log.Panic()
	}
	fmt.Println(publish)

	//stop := time.Since(start)
	//fmt.Println(stop)
}
