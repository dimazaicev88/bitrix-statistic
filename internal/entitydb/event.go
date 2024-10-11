package entitydb

import (
	"github.com/google/uuid"
	"time"
)

type (
	Event struct {
		Uuid            uuid.UUID `ch:"uuid"`
		Event1          string    `ch:"event1"`
		Event2          string    `ch:"event2"`
		Money           float64   `ch:"money"`
		DateEnter       time.Time `ch:"date_enter"`
		DateCleanup     time.Time `ch:"date_cleanup"`
		Sort            uint32    `ch:"sort"`
		Counter         uint32    `ch:"counter"`
		AdvVisible      bool      `ch:"adv_visible"`
		Name            string    `ch:"name"`
		Description     string    `ch:"description"`
		KeepDays        uint32    `ch:"keep_days"`
		DynamicKeepDays uint32    `ch:"dynamic_keep_days"`
		DiagramDefault  bool      `ch:"diagram_default"`
	}

	EventDay struct {
		Uuid      uuid.UUID `ch:"uuid"`
		DateStat  time.Time `ch:"date_stat"`
		DateLast  time.Time `ch:"date_last"`
		EventUuid uuid.UUID `ch:"event_uuid"`
		Money     float64   `ch:"money"`
		Counter   uint32    `ch:"counter"`
	}

	EventList struct {
		Uuid          uuid.UUID `ch:"uuid"`
		EventUuid     uuid.UUID `ch:"event_uuid"`
		Event3        string    `ch:"event3"`
		Money         float64   `ch:"money"`
		DateEnter     time.Time `ch:"date_enter"`
		refererUrl    string    `ch:"referer_url"`
		URL           string    `ch:"url"`
		RedirectUrl   string    `ch:"redirect_url"`
		SessionUuid   uuid.UUID `ch:"session_uuid"`
		GuestUuid     uuid.UUID `ch:"guest_uuid"`
		GuestAdvUuid  uuid.UUID `ch:"guest_adv_uuid"`
		AdvBack       bool      `ch:"adv_back"`
		HitUuid       uuid.UUID `ch:"hit_uuid"`
		CountryId     string    `ch:"country_id"`
		KeepDays      uint32    `ch:"keep_days"`
		Chargeback    bool      `ch:"chargeback"`
		SiteId        string    `ch:"site_id"`
		RefererSiteId string    `ch:"referer_site_id"`
	}
)
