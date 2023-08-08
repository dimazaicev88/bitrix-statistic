package main

import (
	"errors"
	"net/url"
	"regexp"
	"strings"
)

type UrlToSql struct {
	urlValues url.Values
}

/**
api/v1/hit?id=20 +
api/v1/hit?id=gt:20 //>20 +
api/v1/hit?id=gte:20 //>=20+
api/v1/hit?id=lt:20 //<20
api/v1/hit?id=lte:20 //<=20
api/v1/hit?id=[20]or[21]or[50] //20 or 21 or 50
api/v1/hit?id=1..20_or_21_or_50 //(id>1 and id<20) or id>21 or 50<id
api/v1/hit?id=20_and_21_and_50 //20 or 21 or 50
api/v1/hit?id=lt:20&url=~samson%_or_~%guest% // like 'samson%'
api/v1/hit?id=lt:20&url=~"samson%_"%_or_~"%guest%" // like 'samson%'
*/

func NewUrlToSql(strUrl string) (UrlToSql, error) {
	parse, err := url.Parse(strUrl)
	if err != nil {
		return UrlToSql{}, err
	}
	return UrlToSql{urlValues: parse.Query()}, nil
}

func (p UrlToSql) ToSqlIntField(fieldName string) (string, error) {
	var sql strings.Builder

	if !p.urlValues.Has(fieldName) {
		return "", errors.New("field: " + fieldName + " not found")
	}
	rgOnlyNum, _ := regexp.Compile(`^\d+$`)
	rgGt, _ := regexp.Compile(`^gt:\d+$`)
	rgGte, _ := regexp.Compile(`^gte:\d+$`)
	rgLt, _ := regexp.Compile(`^lt:\d+$`)
	rgLte, _ := regexp.Compile(`^lte:\d+$`)
	if rgOnlyNum.MatchString(p.urlValues.Get(fieldName)) {
		sql.WriteString(fieldName)
		sql.WriteString("=")
		sql.WriteString(p.urlValues.Get(fieldName))
	}

	if rgGt.MatchString(p.urlValues.Get(fieldName)) {
		sql.WriteString(fieldName)
		sql.WriteString(">")
		sql.WriteString(strings.Split(p.urlValues.Get(fieldName), ":")[1])
	}

	if rgGte.MatchString(p.urlValues.Get(fieldName)) {
		sql.WriteString(fieldName)
		sql.WriteString(">=")
		sql.WriteString(strings.Split(p.urlValues.Get(fieldName), ":")[1])
	}

	if rgLt.MatchString(p.urlValues.Get(fieldName)) {
		sql.WriteString(fieldName)
		sql.WriteString("<")
		sql.WriteString(strings.Split(p.urlValues.Get(fieldName), ":")[1])
	}

	if rgLte.MatchString(p.urlValues.Get(fieldName)) {
		sql.WriteString(fieldName)
		sql.WriteString("<=")
		sql.WriteString(strings.Split(p.urlValues.Get(fieldName), ":")[1])
	}

	return sql.String(), nil
}

//func main() {
//	//urlToSQl, _ := NewUrlToSql("https://korautotest1.officemag.ru/api/v1/hit?id=gte:20")
//	//value, err := urlToSQl.ToSqlIntField("id")
//	//if err != nil {
//	//	panic(err)
//	//}
//	//
//	//println(value)
//
//	rg, _ := regexp.Compile(`^[gt|le:\[\d+\](or|and)(\[\d+\])]+?$`)
//	rs := rg.MatchString("[lt:10]or[gt:20]or[40]and[50]or[60]or[70]or[80]")
//	fmt.Println(rs)
//
//	//start := time.Now()
//	//rx, _ := regexp.Compile(`id=eq:\d+`)
//	//rx.MatchString("id=eq:200&limit=10&order=asc")
//	////println(b)
//	//elapsed := time.Since(start)
//	//log.Println("time: ", elapsed.Nanoseconds())
//
//	//client := http.Client{}
//	//start := time.Now()
//	//regexp.MatchString(`^id=eq\:\d+$`, "id=eq:200&limit=10&order=asc")
//	//
//	////client.Get("http://127.0.0.1:125/")
//	//stop := time.Since(start)
//	//fmt.Println(stop)
//	//response, err := client.Get("http://127.0.0.1:3000/")
//	//if err != nil {
//	//	log.Fatalln(err)
//	//}
//	//all, err := ioutil.ReadAll(response.Body)
//	//if err != nil {
//	//	log.Fatalln(err)
//	//}
//	//fmt.Println(string(all))
//}
