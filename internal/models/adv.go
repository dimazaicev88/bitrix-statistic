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

func (am Adv) SetAdv(fullRequestUri, openstat, referringSite string) error {

	return nil
}

// FindByByPage Поиск Рекламной компании по странице
func (am Adv) FindByByPage(page, direction string) (entitydb.AdvDb, error) {
	strSql := `
		SELECT t_adv.uuid, t_adv.referer1, t_adv.referer2
		FROM adv t_adv
		INNER JOIN adv_page t_adv_page  ON (t_adv_page.adv_uuid = t_adv.uuid and t_adv_page.type='?')
 		WHERE length(t_adv_page.page) > 0 and t_adv_page.page like ?`

	var adv entitydb.AdvDb
	err := am.chClient.QueryRow(am.ctx, strSql, utils.StringConcat("%", page, "%"), direction).ScanStruct(&adv)
	if err != nil {
		return entitydb.AdvDb{}, nil
	}
	return adv, nil
}

func (am Adv) FindByByDomainSearcher(host string) ([]int, string, string, error) {
	//проверяем поисковики
	sql := ` SELECT A.referer1, A.referer2, S.ADV_ID
			 FROM 	adv A,
			        adv_searcher S,
			        searcher_params P
			 WHERE  S.ADV_ID = A.ID and P.SEARCHER_ID = S.SEARCHER_ID and upper(?) like concat("'%'",upper(P.DOMAIN),"'%'")`

	rows, err := am.storage.DB().Query(sql, host)
	if err != nil {
		return nil, "", "", err
	}

	var listIdAdv []int
	var referer1 string
	var referer2 string
	//for rows.Next() {
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
	return listIdAdv, referer1, referer2, nil
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

func (am Adv) FindById(id int) (entitydb.AdvDb, error) {
	var adv entitydb.AdvDb
	//sql := `-- SELECT 	* FROM adv WHERE  id=?`
	//err := am.storage.DB().Get(&adv, sql, id)
	//if err != nil {
	//	return entitydb.AdvDb{}, err
	//}
	return adv, nil
}
