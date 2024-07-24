package entityjson

import "time"

type StopList struct {
	DateStart       time.Time `json:"dateStart"`                 // Время с которого запись начинает действовать
	DateEnd         time.Time `json:"dateEnd"`                   // Время после которого действие записи заканчивается
	Active          bool      `json:"active,omitempty"`          // "Y" - запись активна; "N" - запись не активна
	SaveStatistic   bool      `json:"saveStatistic,omitempty"`   // "Y" - сохранять статистику по посетителю, попавшему в стоп-лист; "N" - не сохранять подобную статистику
	Ip1             string    `json:"ip1,omitempty"`             // Октет 1 IP адреса
	Ip2             string    `json:"ip2,omitempty"`             // Октет 2 IP адреса
	Ip3             string    `json:"ip3,omitempty"`             // Октет 3 IP адреса
	Ip4             string    `json:"ip4,omitempty"`             // Октет 4 IP адреса
	Mask1           string    `json:"mask1,omitempty"`           // Маска для октета 1 IP адреса
	Mask2           string    `json:"mask2,omitempty"`           // Маска для октета 2 IP адреса
	Mask3           string    `json:"mask3,omitempty"`           // Маска для октета 3 IP адреса
	Mask4           string    `json:"mask4,omitempty"`           // Маска для октета 4 IP адреса
	UserAgent       string    `json:"userAgent,omitempty"`       // UserAgent посетителя
	UserAgentIsNull bool      `json:"userAgentIsNull,omitempty"` // "Y" - UserAgent посетителя не задан (пустой);  "N" - UserAgent посетителя задан (значение по умолчанию)
	UrlTo           string    `json:"urlTo,omitempty"`           // Страница (или ее часть) на которую приходит посетитель
	UrlFrom         string    `json:"urlFrom,omitempty"`         // Ссылающаяся страница (или ее часть), с которой приходит посетитель
	Message         string    `json:"message,omitempty"`         // Текст сообщения которое будет выдано посетителю сайта, в случае его попадания под данную запись стоп-листа
	MessageLid      string    `json:"messageLid,omitempty"`      // Язык сообщения задаваемого в поле MESSAGE]
	UrlRedirect     string    `json:"urlRedirect,omitempty"`     // Страница на которую необходимо перенаправить посетителя после его попадания под данную запись стоп-листа
	Comments        string    `json:"comments,omitempty"`        // Административный комментарий; используется, как правило, для указания причин создания данной записи
	Test            bool      `json:"test,omitempty"`            // "Y" - данная запись является тестовой; "N" - данная запись не тестовая (см. метод CStopList::Check)
	SiteId          string    `json:"siteId,omitempty"`          // ID сайта для которого запись будет действительна; если значение не задано, то это означает что запись действительная для всех сайтов
	Lamp            string    `json:"lamp,omitempty"`            // "green" - запись активна и работает; "red" - запись не активная по каким либо причинам (возможно истек срок действия, либо снят флаг активности)
}
