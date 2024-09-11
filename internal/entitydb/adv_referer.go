package entitydb

import "github.com/google/uuid"

type AdvCompany struct {
	AdvUuid     uuid.UUID `ch:"adv_uuid"`
	Referer1    string    `ch:"referer1"`
	Referer2    string    `ch:"referer2"`
	Referer3    string    `ch:"referer3"`
	LastAdvBack bool      `ch:"last_adv_back"`
}
