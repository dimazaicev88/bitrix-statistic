package entity

import "time"

type EventDb struct {
	Uuid            string    `ch:"uuid"`
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
