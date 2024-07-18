package entitydb

type AdvPageDB struct {
	Uuid    string `ch:"uuid"`
	AdvUuid string `ch:"adv_uuid"`
	Page    string `ch:"page"`
	Type    string `ch:"type"`
}
