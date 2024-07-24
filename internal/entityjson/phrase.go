package entityjson

import (
	"github.com/google/uuid"
	"time"
)

type PhraseGroupP struct {
	Phrase   string  `json:"phrase"`   //Поисковая фраза
	Quantity uint32  `json:"quantity"` //Сколько раз заходили на сайт по данной поисковой фразе
	Percent  float32 `json:"percent"`  //Процент от общего количества заходов по разным поисковым фразам
}

type PhraseGroupS struct {
	SearcherUuid uuid.UUID `json:"searcherUuid"` //ID поисковой системы
	SearcherName string    `json:"searcherName"` //Название поисковой системы
	Quantity     uint32    `json:"quantity"`     //Количество заходов с поисковой системы
	Percent      float32   `json:"percent"`      //Процент заходов с данной поисковой системы
	AverageHits  float32   `json:"averageHits"`  //Среднее количество хитов, производимое посетителями, заходящими с той или иной поисковой системы
}

type Phrase struct {
	Uuid         uuid.UUID `json:"uuid"`         //Uuid записи
	PHRASE       string    `json:"phrase"`       //Поисковая фраза
	DateHit      time.Time `json:"dateHit"`      //время
	SessionUuid  uuid.UUID `json:"sessionUuid"`  //ID сессии
	RefererUuid  uuid.UUID `json:"refererUuid"`  //ID записи из таблицы ссылающихся сайтов (страниц)
	SearcherUuid uuid.UUID `json:"searcherUuid"` //ID поисковой системы
	SearcherName string    `json:"searcherName"` //Название поисковой системы
	UrlTo        string    `json:"urlTo"`        //Страница на которую пришли
	UrlTo404     bool      `json:"urlTo404"`     //True|False флаг 404 ошибки на странице, на которую пришли
	SiteId       string    `json:"siteId"`       //Сайта, на который пришли
}
