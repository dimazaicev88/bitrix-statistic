package main

import (
	"os"
)

func main() {
	dat, err := os.ReadFile("E:\\projects\\bitrix-statistic\\sxgeo\\SxGeo.dat")
	if err != nil {
		panic("Database file not found")
	} else if string(dat[:3]) != `SxG` && dat[3] != 22 && dat[8] != 2 {
		panic("Wrong database format")
	} else if dat[9] != 0 {
		panic("Only UTF-8 version is supported")
	}

	//geo := sypexgeo.New("E:\\projects\\bitrix-statistic\\sxgeo\\SxGeo.dat")
	//info, err := geo.GetCityFull("93.73.35.74")
	//logrus.Fatal(err)
	//logrus.Println(info)

	//fmt.Println(time.Now().Local().Format("2006-01-02"))
}
