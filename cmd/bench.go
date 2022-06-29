package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := http.Client{}
	start := time.Now()
	client.Get("http://127.0.0.1:3000/")
	stop := time.Now().Sub(start)
	fmt.Println(stop)
	//response, err := client.Get("http://127.0.0.1:3000/")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//all, err := ioutil.ReadAll(response.Body)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println(string(all))
}
