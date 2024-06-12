package models

import (
	"bitrix-statistic/internal/session"
	"bitrix-statistic/internal/storage"
	"net/url"
	"regexp"
	"strconv"
)

type AdvModel struct {
	storage     storage.Storage
	optionModel *OptionModel
}

func NewAdvModel(storage storage.Storage, optionModel *OptionModel) *AdvModel {
	return &AdvModel{
		storage:     storage,
		optionModel: optionModel,
	}
}

func (am AdvModel) SetAdv(phpSession *session.Session, fullRequestUri, openstat, referringSite string) error {
	//$err_mess = "File: " . __FILE__ . "<br>Line: ";
	//stat_session_register("SESS_ADV_ID") // ID рекламной кампании
	//$DB = CDatabase::GetModuleConnection('statistic');
	var listAdv []int // массив рекламных кампаний
	var ref1, ref2 string
	var err error
	// если это начало сессии
	if phpSession.KeyExists("SESS_SESSION_ID") == false || phpSession.GetAsInt("SESS_SESSION_ID") <= 0 &&
		phpSession.KeyExists("SESS_ADV_ID") == false || phpSession.GetAsInt("SESS_ADV_ID") <= 0 {

		// проверяем страницу на которую пришел посетитель
		//$page_to = __GetFullRequestUri()
		listAdv, ref1, ref2, err = am.FindByByPage(fullRequestUri, "TO")
		if err != nil {
			return err
		}

		// если посетитель пришел с ссылающегося сайта то
		if len(referringSite) > 0 {
			urpParse, _ := url.Parse(referringSite)

			listAdv, ref1, ref2, err = am.FindByByDomainSearcher(urpParse.Host)
			if err != nil {
				return err
			}

			listAdv, ref1, ref2, err = am.FindByByPage(referringSite, "FROM")
			if err != nil {
				return err
			}
		}

		// если гость пришел с referer1, либо referer2 то
		if phpSession.Get("referer1") != "" || phpSession.Get("referer2") != "" {
			listAdv, ref1, ref2, err = am.FindByReferer(phpSession.Get("referer1"), phpSession.Get("referer2"))
		}
		//Handle Openstat if enabled TODO Доделать
		//if am.optionModel.Get("OPENSTAT_ACTIVE") == "Y" && len(openstat) >0 {
		//	if strings.Contains(openstat,";")== false{
		//		result, err := base64.StdEncoding.DecodeString(openstat)
		//		if err != nil {
		//			return err
		//		}
		//		values:=strings.Split(string(result), ";")
		//	}
		//am.FindByReferer(
		//		trim(str_replace(
		//			array("#service-name#", "#campaign-id#", "#ad-id#", "#source-id#"),
		//	$openstat,
		//		COption::GetOptionString("statistic", "OPENSTAT_R1_TEMPLATE")
		//	)),
		//	trim(str_replace(
		//		array("#service-name#", "#campaign-id#", "#ad-id#", "#source-id#"),
		//	$openstat,
		//		COption::GetOptionString("statistic", "OPENSTAT_R2_TEMPLATE")
		//	)),
		//	$arrADV, $ref1, $ref2
		//	)
		//}

		// если было выявлено более одной рекламной кампании подходящей под условия то
		if len(listAdv) > 1 {
			strSql := `SELECT id, referer1, referer2 FROM adv WHERE id in (?) ORDER BY priority desc, id desc`
			rows, err := am.storage.DB().Query(strSql, listAdv)
			if err != nil {
				return nil
			}
			var referer1 string
			var referer2 string
			for rows.Next() {
				var id int
				err = rows.Scan(&id, &referer1, &referer2)
				if err != nil {
					return nil
				}
				phpSession.Set("SESS_ADV_ID", strconv.Itoa(id))
				phpSession.Set("referer1", referer1)
				phpSession.Set("referer2", referer2)
			}
			err = rows.Err()
			if err != nil {
				return nil
			}
		} else {
			phpSession.Set("SESS_ADV_ID", "")
			phpSession.Set("referer1", ref1)
			phpSession.Set("referer2", ref2)
		}
	}
	if phpSession.GetAsInt("SESS_ADV_ID") > 0 {
		phpSession.Set("SESS_LAST_ADV_ID", phpSession.Get("SESS_ADV_ID"))
		phpSession.Set("SESS_LAST_ADV_ID", strconv.Itoa(phpSession.GetAsInt("SESS_LAST_ADV_ID")))
	}
	return nil
}

func (am AdvModel) FindByByPage(page, cType string) ([]int, string, string, error) {
	strSql := `
		SELECT A.id, A.referer1, A.referer2
		FROM adv A
		INNER JOIN adv_page AP ON (AP.adv_id = A.id and AP.c_type='?')
		WHERE length(AP.page) > 0
		and ? like concat('%', AP.page, '%')`
	rows, err := am.storage.DB().Query(strSql, page, cType)
	if err != nil {
		return nil, "", "", err
	}

	var listIdAdv []int
	var referer1 string
	var referer2 string
	for rows.Next() {
		var id int
		err = rows.Scan(&id, &referer1, &referer2)
		if err != nil {
			return nil, "", "", err
		}
		listIdAdv = append(listIdAdv, id)
	}
	err = rows.Err()
	if err != nil {
		return nil, "", "", err
	}

	return listIdAdv, referer1, referer2, nil
}

func (am AdvModel) FindByByDomainSearcher(host string) ([]int, string, string, error) {
	// проверяем поисковики
	sql := ` SELECT A.referer1, A.referer2, S.ADV_ID
			         FROM 	adv A,
				            adv_searcher S,
				            searcher_params P
			         WHERE
			                S.ADV_ID = A.ID and P.SEARCHER_ID = S.SEARCHER_ID and upper(?) like concat("'%'",upper(P.DOMAIN),"'%'")`

	rows, err := am.storage.DB().Query(sql, host)
	if err != nil {
		return nil, "", "", err
	}

	var listIdAdv []int
	var referer1 string
	var referer2 string
	for rows.Next() {
		var id int
		err = rows.Scan(&id, &referer1, &referer2)
		if err != nil {
			return nil, "", "", err
		}
		listIdAdv = append(listIdAdv, id)
	}
	err = rows.Err()
	if err != nil {
		return nil, "", "", err
	}
	return listIdAdv, referer1, referer2, nil
}

func (am AdvModel) FindByReferer(referer1, referer2 string) ([]int, string, string, error) {
	sql := `
	SELECT 	ID, REFERER1, REFERER2
	FROM adv
	WHERE  REFERER1=? and REFERER2=?`

	found := false
	rows, err := am.storage.DB().Query(sql, referer1, referer2)
	if err != nil {
		return nil, "", "", err
	}

	var listIdAdv []int
	for rows.Next() {
		found = true
		var id int
		err = rows.Scan(&id, &referer1, &referer2)
		if err != nil {
			return nil, "", "", err
		}
		listIdAdv = append(listIdAdv, id)
	}
	err = rows.Err()
	if err != nil {
		return nil, "", "", err
	}
	na := ""
	if !found {
		if am.optionModel.Get("ADV_NA") == "Y" {
			Na1 := am.optionModel.Get("AVD_NA_REFERER1")
			Na2 := am.optionModel.Get("AVD_NA_REFERER2")
			if (Na1 != "" || Na2 != "") && referer1 == Na1 && referer2 == Na2 {
				na = "Y"
			}

		}

		if am.optionModel.Get("ADV_AUTO_CREATE") == "Y" || (na == "Y") {
			var bGoodR bool
			if am.optionModel.Get("REFERER_CHECK") == "Y" {
				bGoodR, err = regexp.MatchString("/^([0-9A-Za-z_:;.,-])*$/", referer1)
				if err != nil {
					return nil, "", "", err
				}
				if bGoodR {
					bGoodR, err = regexp.MatchString("/^([0-9A-Za-z_:;.,-])*$/", referer2)
				}
				if err != nil {
					return nil, "", "", err
				}
			} else {
				bGoodR = true
			}

			if bGoodR {
				err := am.AddAdv(referer1, referer2)
				if err != nil {
					return nil, "", "", err
				}
			}
		}
	}
	return listIdAdv, referer1, referer2, nil
}

func (am AdvModel) AddAdv(referer1 string, referer2 string) error {
	_, err := am.storage.DB().MustExec(`INSERT INTO adv(referer1, referer2, date_first, date_last)
		VALUES (?, ?, now(), now())`, referer1, referer2).LastInsertId()
	if err != nil {
		return err
	}
	return nil
}
