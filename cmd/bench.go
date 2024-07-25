package main

import "log"
import "github.com/barsuk/sxgeo"

type ems struct {
	Id int
}

func main() {

	if _, err := sxgeo.ReadDBToMemory(dbPath); err != nil {
		log.Fatalf("error: cannot read database file: %v", err)
	}

	//fmt.Println(ems{} == ems{Id: 1})

	//fmt.Println(cache.AdvDays())

	//var q interface{} = 1844674407370955161
	//
	//switch q.(type) {
	//
	//case int:
	//case float64:
	//case string:
	//default:
	//	logrus.Panic("unknown type")
	//}

	//ms := make(map[string][]int)

	//col1Data, _ := uuid.NewUUID()
	//
	//fmt.Println(col1Data.String())
	//col1Data.String()

	//cache, _ := bigcache.New(context.Background(), bigcache.DefaultConfig(1*time.Minute))
	//
	//for i := 0; i < 1000000; i++ {
	//	err := cache.Set(strconv.Itoa(i), []byte("valueeeeeeeeeeeeeeeeee"))
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}

	//entry, _ := cache.Get("my-unique-key")
	//fmt.Println(string(entry))
}
