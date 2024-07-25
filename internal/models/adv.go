package models

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/utils"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Adv struct {
	chClient    driver.Conn
	optionModel *Option
	ctx         context.Context
}

func NewAdv(ctx context.Context, chClient driver.Conn, optionModel *Option) *Adv {
	return &Adv{
		chClient:    chClient,
		optionModel: optionModel,
		ctx:         ctx,
	}
}

// FindAdvUuidByByPage Поиск Рекламной компании по странице
func (am Adv) FindAdvUuidByByPage(page, direction string) ([]string, error) {
	strSql := `
		SELECT t_adv.uuid
		FROM adv t_adv
		INNER JOIN adv_page t_adv_page  ON (t_adv_page.adv_uuid = t_adv.uuid and t_adv_page.type=?)
 		WHERE length(t_adv_page.page) > 0 and t_adv_page.page like ?`

	var listAdvUuid []string

	rows, err := am.chClient.Query(am.ctx, strSql, direction, utils.StringConcat("%", page, "%"))
	if err != nil {
		return []string{}, err
	}

	for rows.Next() {
		var advUuid string
		if err := rows.Scan(&advUuid); err != nil {
			return []string{}, err
		}
		listAdvUuid = append(listAdvUuid, advUuid)
	}
	return listAdvUuid, nil
}

func (am Adv) FindByByDomainSearcher(host string) ([]string, error) {
	//проверяем поисковики
	sql := ` SELECT t_adv_searcher.adv_uuid
			FROM adv_searcher t_adv_searcher
					 JOIN searcher_params t_searcher_params ON t_adv_searcher.searcher_uuid = t_searcher_params.searcher_uuid
			WHERE t_searcher_params.domain like ?`

	var listAdvSearcherUuid []string
	rows, err := am.chClient.Query(am.ctx, sql, utils.StringConcat("%", host, "%"))
	if err != nil {
		return []string{}, err
	}

	for rows.Next() {
		var advUuid string
		if err := rows.Scan(&advUuid); err != nil {
			return []string{}, err
		}
		listAdvSearcherUuid = append(listAdvSearcherUuid, advUuid)
	}
	return listAdvSearcherUuid, nil
}

func (am Adv) FindByReferer(referer1, referer2 string) ([]int, string, string, error) {
	//sql := `SELECT 	ID, REFERER1, REFERER2
	//		FROM adv
	//		WHERE  REFERER1=? and REFERER2=?`
	//
	//found := false
	//rows, err := am.storage.DB().Query(sql, referer1, referer2)
	//if err != nil {
	//	return nil, "", "", err
	//}

	var listIdAdv []int
	//for rows.Next() {
	//	found = true
	//	var id int
	//	err = rows.Scan(&id, &referer1, &referer2)
	//	if err != nil {
	//		return nil, "", "", err
	//	}
	//	listIdAdv = append(listIdAdv, id)
	//}
	//err = rows.Err()
	//if err != nil {
	//	return nil, "", "", err
	//}
	//na := ""
	//if !found {
	//	if am.optionModel.Get("ADV_NA") == "Y" {
	//		Na1 := am.optionModel.Get("AVD_NA_REFERER1")
	//		Na2 := am.optionModel.Get("AVD_NA_REFERER2")
	//		if (Na1 != "" || Na2 != "") && referer1 == Na1 && referer2 == Na2 {
	//			na = "Y"
	//		}
	//
	//	}
	//
	//	if am.optionModel.Get("ADV_AUTO_CREATE") == "Y" || (na == "Y") {
	//		var bGoodR bool
	//		if am.optionModel.Get("REFERER_CHECK") == "Y" {
	//			bGoodR, err = regexp.MatchString("/^([0-9A-Za-z_:;.,-])*$/", referer1)
	//			if err != nil {
	//				return nil, "", "", err
	//			}
	//			if bGoodR {
	//				bGoodR, err = regexp.MatchString("/^([0-9A-Za-z_:;.,-])*$/", referer2)
	//			}
	//			if err != nil {
	//				return nil, "", "", err
	//			}
	//		} else {
	//			bGoodR = true
	//		}
	//
	//		if bGoodR {
	//			err := am.AddAdv(referer1, referer2)
	//			if err != nil {
	//				return nil, "", "", err
	//			}
	//		}
	//	}
	//}
	return listIdAdv, referer1, referer2, nil
}

func (am Adv) AddAdv(referer1 string, referer2 string) error {
	//_, err := am.storage.DB().MustExec(`INSERT INTO adv(referer1, referer2, date_first, date_last)
	//	VALUES (?, ?, now(), now())`, referer1, referer2).LastInsertId()
	//if err != nil {
	//	return err
	//}
	return nil
}

func (am Adv) FindByUuid(uuid string) (entitydb.AdvDb, error) {
	var adv entitydb.AdvDb
	sql := `SELECT 	* FROM adv WHERE  uuid=?`
	err := am.chClient.QueryRow(am.ctx, sql, uuid).ScanStruct(&adv)
	if err != nil {
		return entitydb.AdvDb{}, err
	}
	return adv, nil
}

func (am Adv) DeleteByUuid(uuid string) error {
	if err := am.chClient.Exec(am.ctx, `DELETE FROM adv WHERE uuid=?`, uuid); err != nil {
		return err
	}
	return nil
}

func (am Adv) FindRefererByListAdv(listAdv []string) (entitydb.AdvReferer, error) {
	var adv entitydb.AdvReferer
	sql := `SELECT 	referer1,referer2 FROM adv WHERE  uuid IN (?) ORDER BY priority,date_create DESC LIMIT 1`
	err := am.chClient.QueryRow(am.ctx, sql, listAdv).ScanStruct(&adv)
	if err != nil {
		return entitydb.AdvReferer{}, err
	}
	return adv, nil
}
