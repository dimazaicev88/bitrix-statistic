package entityjson

type Country struct {
	Id            string `json:"id"`                      //Двух символьный идентификатор страны
	ShortName     string `json:"shortName,omitempty"`     //Трех символьный идентификатор страны
	Name          string `json:"name,omitempty"`          //Название страны
	Sessions      uint32 `json:"sessions,omitempty"`      //Суммарное кол-во сессий
	NewGuests     uint32 `json:"newGuests,omitempty"`     //Суммарное кол-во новых посетителей
	Hits          uint32 `json:"hits,omitempty"`          //Суммарное кол-во хитов
	Events        uint32 `json:"events,omitempty"`        //Суммарное кол-во событий
	ReferenceUuid string `json:"referenceUuid,omitempty"` //параметр REFERENCE_ID для использования в функции SelectBox или SelectBoxM
	Reference     string `json:"reference,omitempty"`     //параметр REFERENCE для использования в функции SelectBox или SelectBoxM
}
