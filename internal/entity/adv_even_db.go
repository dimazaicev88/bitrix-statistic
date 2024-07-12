package entity

type AdvEvenDB struct {
	Uuid        string  `ch:"uuid"`
	AdvUuid     string  `ch:"adv_uuid"`
	EventUuid   string  `ch:"event_uuid"`
	Counter     uint32  `ch:"counter"`
	CounterBack uint32  `ch:"counter_back"`
	Money       float64 `ch:"money"`
	MoneyBack   float64 `ch:"money_back"`
}
