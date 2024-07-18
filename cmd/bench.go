package main

import "github.com/sirupsen/logrus"

func main() {

	var q interface{} = 1844674407370955161

	switch q.(type) {

	case int:
	case float64:
	case string:
	default:
		logrus.Panic("unknown type")
	}

	//ms := make(map[string][]int)

	//col1Data, _ := uuid.NewUUID()
	//
	//fmt.Println(col1Data.String())
	//col1Data.String()

	//optioncache, _ := bigcache.New(context.Background(), bigcache.DefaultConfig(1*time.Minute))
	//
	//for i := 0; i < 1000000; i++ {
	//	err := optioncache.Set(strconv.Itoa(i), []byte("valueeeeeeeeeeeeeeeeee"))
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}

	//entry, _ := optioncache.Get("my-unique-key")
	//fmt.Println(string(entry))
}
