package models

import (
	"bitrix-statistic/internal/session"
	"bitrix-statistic/internal/storage"
	"encoding/json"
)

type SessionDataModel struct {
	storage storage.Storage
}

func NewSessionDataModel(storage storage.Storage) *SessionDataModel {
	return &SessionDataModel{storage: storage}
}

func (sm SessionDataModel) FindSessionByGuestMd5(guestMd5 string, sessionGcMaxLifeTime string) (session.Session, error) {
	row := sm.storage.DB().QueryRow(
		`SELECT session_data FROM session_data
                    WHERE guest_md5=? and date_last > DATE_ADD(now(), INTERVAL-? SECOND) 
                    LIMIT 1`,
		guestMd5, sessionGcMaxLifeTime,
	)
	var sessionDb string
	var sessionData session.Session
	err := row.Scan(&sessionData)
	if err != nil {
		return session.Session{}, err
	}
	if err := json.Unmarshal([]byte(sessionDb), &sessionData); err != nil {
		return session.Session{}, err
	}
	return sessionData, nil
}
