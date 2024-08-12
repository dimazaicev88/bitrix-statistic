package models

import (
	"bitrix-statistic/internal/entitydb"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type PathAdv struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewPathAdv(ctx context.Context, chClient driver.Conn) *PathAdv {
	return &PathAdv{
		ctx:      ctx,
		chClient: chClient,
	}
}

func (pa PathAdv) Add(pathAdv entitydb.PathAdv) error {
	return pa.chClient.Exec(pa.ctx,
		`INSERT INTO path_adv (uuid, adv_uuid, path_id, date_stat, counter, counter_back, counter_full_path, counter_full_path_back, steps, sign, version)
			VALUES (generateUUIDv7(),?,?,curdate(),?,?,?,?,?,?,?)`,
		pathAdv.AdvUuid, pathAdv.Counter, pathAdv.CounterBack, pathAdv.CounterFullPath, pathAdv.CounterFullPathBack, pathAdv.Steps, pathAdv.Sign, pathAdv.Version,
	)
}

func (pa PathAdv) FindByPathId(id int32, dateStat string) (entitydb.PathAdv, error) {
	var pathAdv entitydb.PathAdv
	err := pa.chClient.QueryRow(pa.ctx, `SELECT * FROM path_adv WHERE path_id = ? and date_stat=?`, id, dateStat).Scan(&pathAdv)
	if err != nil {
		return entitydb.PathAdv{}, err
	}
	return pathAdv, nil
}

func (pa PathAdv) Update(oldValue entitydb.PathAdv, newValue entitydb.PathAdv) error {
	err := pa.chClient.Exec(pa.ctx,
		`INSERT INTO path_adv (uuid, adv_uuid, path_id, date_stat, counter, counter_back, counter_full_path, counter_full_path_back, steps, sign, version)
			VALUES (generateUUIDv7(),?,?,curdate(),?,?,?,?,?,?,?)`,
		oldValue.AdvUuid, oldValue.Counter, oldValue.CounterBack, oldValue.CounterFullPath, oldValue.CounterFullPathBack, oldValue.Steps, oldValue.Sign*-1, oldValue.Version,
	)
	if err != nil {
		return err
	}

	err = pa.chClient.Exec(pa.ctx,
		`INSERT INTO path_adv (uuid, adv_uuid, path_id, date_stat, counter, counter_back, counter_full_path, counter_full_path_back, steps, sign, version)
			VALUES (generateUUIDv7(),?,?,curdate(),?,?,?,?,?,?,?)`,
		oldValue.AdvUuid, oldValue.Counter, oldValue.CounterBack, oldValue.CounterFullPath, oldValue.CounterFullPathBack, oldValue.Steps, 1, oldValue.Version+1,
	)
	if err != nil {
		return err
	}
	return nil
}
