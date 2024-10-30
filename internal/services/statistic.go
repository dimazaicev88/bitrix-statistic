package services

import (
	"bitrix-statistic/internal/dto"
	"bitrix-statistic/internal/entitydb"
	"github.com/google/uuid"
	"github.com/maypok86/otter"
	"github.com/sirupsen/logrus"
	_ "net/netip"
	"time"
)

type Statistic struct {
	guestService   *GuestService
	sessionService *SessionService
	hitService     *HitService
	otterCache     otter.Cache[string, entitydb.Session]
}

func NewStatistic() *Statistic {
	otterCache, err := otter.MustBuilder[string, entitydb.Session](15000).
		CollectStats().
		WithTTL(time.Minute * 15).
		Build()

	if err != nil {
		logrus.Fatal(err)
	}

	return &Statistic{
		otterCache: otterCache,
	}
}

func (stat *Statistic) SetHitService(hitService *HitService) {
	stat.hitService = hitService
}

func (stat *Statistic) SetGuestService(guestService *GuestService) {
	stat.guestService = guestService
}

func (stat *Statistic) SetSessionService(sessionService *SessionService) {
	stat.sessionService = sessionService
}

func (stat *Statistic) Add(statData dto.UserData) error {
	var sessionDb entitydb.Session
	var guestDb entitydb.Guest
	var err error
	isNewGuest := true
	var hitUuid = uuid.New()
	var sessionUuid = uuid.New()

	//--------------------------- Guest ------------------------------------
	guestDb, err = stat.guestService.FindByUuid(statData.GuestUuid)
	if err != nil {
		return err
	}

	//Гость не найден, добавляем гостя
	if guestDb == (entitydb.Guest{}) {
		guestDb, err = stat.guestService.Add(statData)
		if err != nil {
			return err
		}
	} else {
		isNewGuest = false
	}

	//--------------------------- Sessions ------------------------------------
	sessionDb, err = stat.sessionService.FindByPHPSessionId(statData.PHPSessionId)
	if err != nil {
		return err
	}

	//Если сессия новая, добавляем.
	if sessionDb == (entitydb.Session{}) {
		sessionDb, err = stat.sessionService.Add(sessionUuid, guestDb.Uuid, hitUuid, statData.PHPSessionId)
		if err != nil {
			return err
		}
	}

	//------------------------------- Hits ---------------------------------
	if _, err = stat.hitService.Add(hitUuid, isNewGuest, sessionUuid, statData); err != nil {
		return err
	}

	return nil
}

func (stat *Statistic) ClearCache() {
	stat.otterCache.Close()
}
