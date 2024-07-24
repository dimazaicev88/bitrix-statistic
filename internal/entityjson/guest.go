package entityjson

import (
	"github.com/google/uuid"
	"time"
)

type Guest struct {
	Uuid            uuid.UUID `json:"uuid,omitempty"`            //ID посетителя
	Favorites       bool      `json:"Favorites,omitempty"`       //Y|N] флаг добавления сайта в "Избранное"
	Events          uint32    `json:"events,omitempty"`          //Количество событий сгенерированных данным посетителей
	Sessions        uint32    `json:"sessions,omitempty"`        //Количество сессий данного посетителя
	Hits            uint32    `json:"hits,omitempty"`            //Количество хитов данного посетителя
	FirstSessionId  uuid.UUID `json:"firstSessionId,omitempty"`  //ID сессии первого захода на сайт
	FirstDate       time.Time `json:"firstDate"`                 //Время первого захода на сайт
	FirstUrlFrom    string    `json:"firstUrlFrom,omitempty"`    //Адрес страницы с которой посетитель впервые пришел на сайт
	FirstUrlTo      string    `json:"firstUrlTo,omitempty"`      //Адрес страницы сайта на которую посетитель впервые пришел
	FirstUrlTo404   bool      `json:"firstUrlTo404,omitempty"`   //Флаг 404 ошибки (страница не существует) на странице сайта на которую посетитель впервые пришел
	FirstSiteId     string    `json:"firstSiteId,omitempty"`     //ID сайта на который посетитель впервые пришел
	FirstAdvUuid    uuid.UUID `json:"firstAdvUuid,omitempty"`    //ID рекламной кампании по которой посетитель впервые пришел на сайт
	FirstReferer1   string    `json:"firstReferer1,omitempty"`   //идентификатор referer1 рекламной кампании FIRST_ADV_ID
	FirstReferer2   string    `json:"firstReferer2,omitempty"`   //идентификатор referer2 рекламной кампании FIRST_ADV_ID
	FirstReferer3   string    `json:"firstReferer3,omitempty"`   //дополнительный параметр referer3 рекламной кампании FIRST_ADV_ID
	LastSessionUuid uuid.UUID `json:"lastSessionUuid,omitempty"` //ID сессии последнего захода на сайт
	LastDate        time.Time `json:"lastDate"`                  //Время последнего захода на сайт
	LastUserId      uint32    `json:"lastUserId,omitempty"`      //ID пользователя
	LastUserAuth    bool      `json:"lastUserAuth,omitempty"`    //True|False был ли авторизован посетитель в последнем заходе на сайт
	LastUrlLast     string    `json:"lastUrlLast,omitempty"`     //Адрес последней страницы на которую зашел посетитель
	LastUrlLast404  bool      `json:"lastUrlLast404,omitempty"`  //Флаг 404 ошибки (страница не существует) на последней странице сайта на которую зашел посетитель
	LastUserAgent   string    `json:"lastUserAgent,omitempty"`   //UserAgent посетителя в последнем заходе
	LastIp          string    `json:"lastIp,omitempty"`          //IP адрес посетителя сайта в последнем заходе
	LastLanguage    string    `json:"lastLanguage,omitempty"`    //Языки установленные в настройках браузера посетителя в последнем заходе
	LastAdvUuid     uuid.UUID `json:"lastAdvUuid,omitempty"`     //ID рекламной кампании по которой посетитель пришел на сайт в последнем заходе
	LastAdvBack     bool      `json:"lastAdvBack,omitempty"`     //True|False флаг того был ли это возврат (Y) или прямой заход (N) по рекламной кампании LAST_ADV_ID
	LastReferer1    string    `json:"lastReferer1,omitempty"`    //идентификатор referer1 рекламной кампании LAST_ADV_ID
	LastReferer2    string    `json:"lastReferer2,omitempty"`    //идентификатор referer2 рекламной кампании LAST_ADV_ID
	LastReferer3    string    `json:"lastReferer3,omitempty"`    //дополнительный параметр referer3 рекламной кампании LAST_ADV_ID
	LastSiteId      string    `json:"lastSiteId,omitempty"`      //ID сайта последнего захода
	LastCountryId   string    `json:"lastCountryId,omitempty"`   //ID страны посетителя в последнем заходе
	LastCountryName string    `json:"lastCountryName,omitempty"` //Название страны посетителя в последнем заходе (если установлено filter"COUNTRY_ID"])
}
