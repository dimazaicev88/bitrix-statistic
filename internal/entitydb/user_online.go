package entitydb

import (
	"github.com/google/uuid"
	"time"
)

type UserOnline struct {
	sessionUuid  uuid.UUID //id сессии
	hits         uint32    //Количество хитов сессии
	lastUserId   uint32    //Id пользователя, под которым последний раз был авторизован посетитель
	userAuth     bool      //Флаг "авторизован ли посетитель в данной сессии" (y - да; n - нет)
	stopListUuid uuid.UUID //Id записи стоп-листа, под которую попал посетитель (если это имело место)
	guestUuid    uuid.UUID //Id посетителя
	newGuest     bool      //Флаг "новый посетитель" (y - новый; n - вернувшийся)
	favorites    bool      //Флаг "добавлял ли посетитель сайт в "избранное" в данной сессии
	countryId    string    //Id страны посетителя
	countryName  string    //Название страны посетителя
	advUuid      uuid.UUID //Id рекламной кампании
	advBack      bool      //Флаг прямого захода (n) или возврата (y) по рекламной кампании
	referer1     string    //Идентификатор referer1 рекламной кампании
	referer2     string    //Идентификатор referer2 рекламной кампании
	referer3     string    //Дополнительный параметр рекламной кампании
	firstUrlFrom string    //Ссылающаяся страница, с которой посетитель впервые пришел на сайт
	urlFrom      string
	firstSiteId  string
	urlLast      string    //Страница последнего хита сессии
	urlLast404   bool      //Флаг 404 ошибки на последней странице сессии (y - да; n - нет)
	lastSiteId   string    //Id сайта на последнем хите сессии
	ipLast       string    //Ip адрес посетителя на последнем хите сессии (в виде: xxx.xxx.xxx.xxx)
	dateLast     time.Time //Время последнего хита
	sessionTime  uint64    //Разница во времени между первым и последним хитом сессии (сек.)
}
