package models

import (
	"bitrix-statistic/internal/entity"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/sirupsen/logrus"
)

type StatisticModel struct {
	ctx         context.Context
	chClient    driver.Conn
	guestModel  *GuestModel
	optionModel *OptionModel
	advModel    *AdvModel
	logger      logrus.Logger
}

func NewStatisticModel(ctx context.Context, chClient driver.Conn) *StatisticModel {
	return &StatisticModel{
		ctx:         ctx,
		chClient:    chClient,
		guestModel:  NewGuestModel(ctx, chClient),
		optionModel: NewOptionModel(ctx, chClient),
		advModel:    NewAdvModel(ctx, chClient, NewOptionModel(ctx, chClient)),
	}
}

func (stm *StatisticModel) Add(data entity.StatData) error {
	guestDb, err := stm.guestModel.FindByHash(data.GuestHash)
	if err != nil {
		stm.logger.Error(err)
		return err
	}

	if len(guestDb) == 0 { //Если пользователь не найден, считаем его новым
		err := stm.guestModel.AddGuest(data)
		if err != nil {
			stm.logger.Error(err)
			return err
		}
	}

	return nil
}

func (stm *StatisticModel) SetNewDay() {

}
