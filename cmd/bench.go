package main

import (
	"context"
	"log"
	"strconv"
	"time"
)
import "github.com/allegro/bigcache/v3"

func main() {

	cache, _ := bigcache.New(context.Background(), bigcache.DefaultConfig(1*time.Minute))

	for i := 0; i < 1000000; i++ {
		err := cache.Set(strconv.Itoa(i), []byte("valueeeeeeeeeeeeeeeeee"))
		if err != nil {
			log.Fatal(err)
		}
	}

	//entry, _ := cache.Get("my-unique-key")
	//fmt.Println(string(entry))
}
