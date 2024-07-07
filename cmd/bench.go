package main

import "fmt"

func main() {

	var q interface{}
	q = "qwe"

	w, ok := q.(int)
	fmt.Println(w, ok)

	//c := gocent.New(gocent.Config{
	//	Addr: "https://centrifugo.intsite.org/api",
	//	Key:  "",
	//})
	//
	//ch := "cicd/portal/chat"
	//ctx := context.Background()
	////start := time.Now()
	//publish, err := c.Publish(ctx, ch, []byte(`{"user": "zabbix","msg": "test message","chatId": 1924358}`))
	//if err != nil {
	//	log.Panic()
	//}
	//fmt.Println(publish)

	//stop := time.Since(start)
	//fmt.Println(stop)
}
