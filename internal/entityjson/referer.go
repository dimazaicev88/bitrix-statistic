package entityjson

import (
	"github.com/google/uuid"
	"time"
)

type RefererGroupU struct {
	UrlFrom     string  `json:"urlFrom,omitempty"`     //Ссылающаяся страница
	Quantity    uint32  `json:"quantity,omitempty"`    //Сколько раз зашли с данной ссылающейся страницы
	Percent     float32 `json:"percent,omitempty"`     //Процент заходов с данной ссылающейся страницы, относительно всех заходов со всех ссылающихся сайтов
	AverageHits float32 `json:"averageHits,omitempty"` //Среднее количество хитов, производимое посетителями, заходящими с данной ссылающейся страницы
}

type RefererGroupS struct {
	UrlFrom     string  `json:"urlFrom,omitempty"`     //Ссылающийся домен
	Quantity    uint32  `json:"quantity,omitempty"`    //Сколько раз зашли с данного ссылающегося домена
	Percent     uint32  `json:"percent,omitempty"`     //Процент заходов с данного ссылающегося домена, относительно всех заходов со всех ссылающихся сайтов
	AverageHits float32 `json:"averageHits,omitempty"` //Среднее количество хитов, производимое посетителями, заходящими с данного ссылающегося домена
}

type Referer struct {
	Uuid        uuid.UUID `json:"uuid,omitempty"`        //ID записи
	UrlFrom     string    `json:"urlFrom,omitempty"`     //Ссылающаяся страница
	DateHit     time.Time `json:"dateHit"`               //дата
	SessionUuid uuid.UUID `json:"sessionUuid,omitempty"` //ID сессии
	UrlTo       string    `json:"urlTo,omitempty"`       //Страница, на которую пришли
	UrlTo404    bool      `json:"urlTo404,omitempty"`    //True|False флаг 404 ошибки на странице на которую пришли
	SiteId      string    `json:"siteId,omitempty"`      //ID сайта на который пришли
}
