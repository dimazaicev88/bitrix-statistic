package entitydb

import (
	"github.com/google/uuid"
	"time"
)

type (
	Adv struct {
		Uuid        uuid.UUID `ch:"uuid"`
		Referer1    string    `ch:"referer1"`
		Referer2    string    `ch:"referer2"`
		Cost        float64   `ch:"cost"`
		DateCreated time.Time `ch:"date_create"`
		EventsView  string    `ch:"events_view"`
		Description string    `ch:"description"`
		Priority    uint32    `ch:"priority"`
		//Referer3    string    `ch:"referer3"`
	}
	AdvStat struct {
		AdvUuid       uuid.UUID `ch:"adv_uuid"`
		Guests        uint32    `ch:"guests"`
		NewGuests     uint32    `ch:"new_guests"`
		Favorites     uint32    `ch:"favorites"`
		Hosts         uint32    `ch:"hosts"`
		Sessions      uint32    `ch:"sessions"`
		Hits          uint32    `ch:"hits"`
		GuestsBack    uint32    `ch:"guests_back"`
		FavoritesBack uint32    `ch:"favorites_back"`
		HostsBack     uint32    `ch:"hosts_back"`
		SessionsBack  uint32    `ch:"sessions_back"`
		HitsBack      uint32    `ch:"hits_back"`
	}

	AdvDynamicResult struct {
		AdvDynamic []AdvDynamic
		AdvMaxMin  *AdvMaxMin
	}

	AdvDynamic struct {
		DateStat      time.Time `ch:"dateStat,omitempty"`      //Дата
		MonthDay      uint8     `ch:"day"`                     //День (1-31)
		Month         uint8     `ch:"month,omitempty"`         //Месяц (1-12)
		Year          uint16    `ch:"year,omitempty"`          //Год
		Guests        uint32    `ch:"guests,omitempty"`        //Посетителей на прямом заходе
		NewGuests     uint32    `ch:"newGuests,omitempty"`     //Новых посетителей на прямом заходе
		Favorites     uint32    `ch:"favorites,omitempty"`     //Посетителей, добавивших сайт в "избранное" на прямом заходе
		Hosts         uint32    `ch:"hosts,omitempty"`         //Хостов на прямом заходе
		Sessions      uint32    `ch:"sessions,omitempty"`      //Сессий на прямом заходе
		Hits          uint32    `ch:"hits,omitempty"`          //Хитов на прямом заходе
		GuestsBack    uint32    `ch:"guestsBack,omitempty"`    //Посетителей на возврате
		FavoritesBack uint32    `ch:"favoritesBack,omitempty"` //Посетителей, добавивших сайт в "избранное" на возврате
		HostsBack     uint32    `ch:"hostsBack,omitempty"`     //Хостов на возврате
		SessionsBack  uint32    `ch:"sessionsBack,omitempty"`  //Сессий на возврате
		HitsBack      uint32    `ch:"hitsBack,omitempty"`      //Хитов на возврате

	}

	AdvMaxMin struct {
		DateFirst time.Time `ch:"dateFirst"`          //Минимальная дата
		MinDay    uint8     `ch:"minDay,omitempty"`   //День минимальной даты (1-31)
		MinMonth  uint8     `ch:"minMonth,omitempty"` //Месяц минимальной даты (1-12)
		MinYear   uint16    `ch:"minYear,omitempty"`  //Год минимальной даты
		DateLast  time.Time `ch:"dateLast"`           //Максимальная дата
		MaxDay    uint8     `ch:"maxDay,omitempty"`   //День максимальной даты (1-31)
		MaxMonth  uint8     `ch:"maxMonth,omitempty"` //Месяц максимальной даты (1-12)
		MaxYear   uint16    `ch:"maxYear,omitempty"`  //Год максимальной даты
	}

	AdvDay struct {
		Uuid          uuid.UUID `ch:"uuid"`
		AdvUuid       string    `ch:"adv_uuid"`
		DateStat      time.Time `ch:"date_stat"`
		Guests        uint32    `ch:"guests"`
		GuestsDay     uint32    `ch:"guests_day"`
		NewGuests     uint32    `ch:"new_guests"`
		Favorites     uint32    `ch:"favorites"`
		Hosts         uint32    `ch:"hosts"`
		HostsDay      uint32    `ch:"hosts_day"`
		Sessions      uint32    `ch:"sessions"`
		Hits          uint32    `ch:"hits"`
		GuestsBack    uint32    `ch:"guests_back"`
		GuestsDayBack uint32    `ch:"guests_day_back"`
		FavoritesBack uint32    `ch:"favorites_back"`
		HostsBack     uint32    `ch:"hosts_back"`
		HostsDayBack  uint32    `ch:"hosts_day_back"`
		SessionsBack  uint32    `ch:"sessions_back"`
		HitsBack      uint32    `ch:"hits_back"`
	}

	AdvEvent struct {
		Uuid        uuid.UUID `ch:"uuid"`
		AdvUuid     uuid.UUID `ch:"adv_uuid"`
		EventUuid   string    `ch:"event_uuid"`
		Counter     uint32    `ch:"counter"`
		CounterBack uint32    `ch:"counter_back"`
		Money       float64   `ch:"money"`
		MoneyBack   float64   `ch:"money_back"`
	}

	AdvEventDay struct {
		Uuid        uuid.UUID `ch:"uuid"`
		AdvUuid     uuid.UUID `ch:"adv_uuid"`
		EventUuid   uuid.UUID `ch:"event_uuid"`
		DateStat    time.Time `ch:"date_stat"`
		Counter     uint32    `ch:"counter"`
		CounterBack uint32    `ch:"counter_back"`
		Money       float64   `ch:"money"`
		MoneyBack   float64   `ch:"money_back"`
	}

	AdvGuest struct {
		Uuid         uuid.UUID `ch:"uuid"`
		AdvUuid      uuid.UUID `ch:"adv_uuid"`
		Back         bool      `ch:"back"`
		GuestUuid    string    `ch:"guest_uuid"`
		DateGuestHit time.Time `ch:"date_guest_hit"`
		DateHostHit  time.Time `ch:"date_host_hit"`
		SessionUuid  string    `ch:"session_uuid"`
		Ip           string    `ch:"ip"`
	}

	AdvPage struct {
		Uuid    uuid.UUID `ch:"uuid"`
		AdvUuid uuid.UUID `ch:"adv_uuid"`
		Page    string    `ch:"page"`
		Type    string    `ch:"type"`
	}

	AdvCompany struct {
		AdvUuid     uuid.UUID `ch:"adv_uuid"`
		Referer1    string    `ch:"referer1"`
		Referer2    string    `ch:"referer2"`
		Referer3    string    `ch:"referer3"`
		LastAdvBack bool      `ch:"last_adv_back"`
	}

	// AdvSearcher TODO Возможно нужно удалить
	AdvSearcher struct {
		Uuid         uuid.UUID `ch:"uuid"`
		AdvUuid      uuid.UUID `ch:"adv_uuid"`
		SearcherUuid string    `ch:"searcher_uuid"`
	}
)
