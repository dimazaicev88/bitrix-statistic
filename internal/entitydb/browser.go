package entitydb

import "github.com/google/uuid"

type Browser struct {
	Uuid      uuid.UUID `ch:"uuid"`
	UserAgent string    `ch:"user_agent"`
}
