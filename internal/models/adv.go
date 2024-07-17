package models

import (
	"bitrix-statistic/internal/entity"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Adv struct {
	chClient    driver.Conn
	optionModel *Option
}

func NewAdv(ctx context.Context, chClient driver.Conn, optionModel *Option) *Adv {
	return &Adv{
		chClient:    chClient,
		optionModel: optionModel,
	}
}

func (am Adv) SetAdv(fullRequestUri, openstat, referringSite string) error {

	return nil
}

func (am Adv) FindByByPage(page, cType string) ([]int, string, string, error) {
	//strSql := `
	//	SELECT A.id, A.referer1, A.referer2
	//	FROM adv A
	//	INNER JOIN adv_page AP ON (AP.adv_id = A.id and AP.c_type='?')
	//	WHERE length(AP.page) > 0 and ? like concat('%', AP.page, '%')`
	//rows, err := am.storage.DB().Query(strSql, page, cType)
	//if err != nil {
	//	return nil, "", "", err
	//}

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

func (am Adv) FindByByDomainSearcher(host string) ([]int, string, string, error) {
	// проверяем поисковики
	//sql := ` SELECT A.referer1, A.referer2, S.ADV_ID
	//		 FROM 	adv A,
	//		        adv_searcher S,
	//		        searcher_params P
	//		 WHERE  S.ADV_ID = A.ID and P.SEARCHER_ID = S.SEARCHER_ID and upper(?) like concat("'%'",upper(P.DOMAIN),"'%'")`
	//
	//rows, err := am.storage.DB().Query(sql, host)
	//if err != nil {
	//	return nil, "", "", err
	//}

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

func (am Adv) FindById(id int) (entity.AdvDb, error) {
	var adv entity.AdvDb
	//sql := `-- SELECT 	* FROM adv WHERE  id=?`
	//err := am.storage.DB().Get(&adv, sql, id)
	//if err != nil {
	//	return entity.AdvDb{}, err
	//}
	return adv, nil
}
