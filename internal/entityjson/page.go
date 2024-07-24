package entityjson

type Page struct {
	Url          string `json:"url"`          //Страница (раздел)
	Dir          bool   `json:"dir"`          //Y|N] флаг "страница(N) или каталог(Y)"
	SiteId       string `json:"siteId"`       //ID сайта
	Url404       bool   `json:"url404"`       // Флаг 404 ошибки для страницы
	Counter      uint32 `json:"counter"`      // Счетчик хитов на данной странице (в данном каталоге) (только если counter_type<>ENTER_COUNTER и counter_type<>EXIT_COUNTER)
	EnterCounter uint32 `json:"enterCounter"` // Счетчик сколько раз данная страница (каталог) 	являлась точкой входа (только при установленном	counter_type=ENTER_COUNTER)
	ExitCounter  uint32 `json:"exitCounter"`  // Счетчик сколько раз данная страница (каталог)	являлась точкой выхода (только при установленном counter_type=EXIT_COUNTER)
}
