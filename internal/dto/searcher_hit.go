package dto

import (
	"time"
)

type SearcherHit struct {
	Uuid         string    `json:"uuid"`                   //ID хита
	DateHit      time.Time `json:"dateHit"`                //Дата хита
	SearcherUuid string    `json:"searcherUuid,omitempty"` //ID поисковой системы
	SearcherName string    `json:"searcherName,omitempty"` //Название поисковой системы
	Url          string    `json:"url,omitempty"`          //Адрес проиндексированной страницы
	Url404       bool      `json:"url404,omitempty"`       //True|False флаг 404 ошибки на проиндексированной странице
	IP           string    `json:"IP,omitempty"`           //IP адрес поисковой системы
	UserAgent    string    `json:"userAgent,omitempty"`    //UserAgent поисковой системы
	HitKeepDays  uint32    `json:"hitKeepDays,omitempty"`  //Индивидуально количество дней, отводимое для хранения хитов поисковой системы
	SiteId       string    `json:"siteId,omitempty"`       // ID сайта
}
