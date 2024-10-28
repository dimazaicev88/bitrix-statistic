package dto

import "time"

type (
	Page struct {
		Url          string `json:"url"`          //Страница (раздел)
		Dir          bool   `json:"dir"`          //Y|N] флаг "страница(N) или каталог(Y)"
		SiteId       string `json:"siteId"`       //ID сайта
		Url404       bool   `json:"url404"`       // Флаг 404 ошибки для страницы
		Counter      uint32 `json:"counter"`      // Счетчик хитов на данной странице (в данном каталоге) (только если counter_type<>ENTER_COUNTER и counter_type<>EXIT_COUNTER)
		EnterCounter uint32 `json:"enterCounter"` // Счетчик сколько раз данная страница (каталог) 	являлась точкой входа (только при установленном	counter_type=ENTER_COUNTER)
		ExitCounter  uint32 `json:"exitCounter"`  // Счетчик сколько раз данная страница (каталог)	являлась точкой выхода (только при установленном counter_type=EXIT_COUNTER)
	}

	PageDynamicList struct {
		DateStat     time.Time `json:"dateStat"`               //дата
		Day          uint8     `json:"day,omitempty"`          // Номер дня (1-31)
		Month        uint8     `json:"month,omitempty"`        // Номер месяца (1-12)
		Year         uint8     `json:"year,omitempty"`         // Год
		Counter      uint32    `json:"counter,omitempty"`      //Кол-во хитов на странице url
		EnterCounter uint32    `json:"enterCounter,omitempty"` //Сколько раз данная страница была точкой входа
		ExitCounter  uint32    `json:"exitCounter,omitempty"`  //Сколько раз данная страница была точкой выхода
	}
)
