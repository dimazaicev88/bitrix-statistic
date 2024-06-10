package models

import (
	"bitrix-statistic/internal/session"
	"bitrix-statistic/internal/storage"
	"encoding/base64"
	"net/url"
	"strings"
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
			// выберем рекламную кампанию по наивысшему приоритету (либо по наивысшему ID)
			$str = implode(",", $arrADV)
			$strSql = "SELECT ID, REFERER1, REFERER2 FROM b_stat_adv WHERE ID in ($str) ORDER BY PRIORITY desc, ID desc"
			$z = $DB- > Query($strSql, false, $err_mess.__LINE__)
			$zr = $z- > Fetch()
			$_SESSION["SESS_ADV_ID"] = intval($zr["ID"])
			$_SESSION["referer1"] = $zr["REFERER1"]
			$_SESSION["referer2"] = $zr["REFERER2"]
		} else {
			$value = reset($arrADV)
			$_SESSION["SESS_ADV_ID"] = intval($value)
			$_SESSION["referer1"] = $ref1
			$_SESSION["referer2"] = $ref2
		}
	}
	if intval($_SESSION["SESS_ADV_ID"]) > 0) $_SESSION["SESS_LAST_ADV_ID"] = $_SESSION["SESS_ADV_ID"]
	$_SESSION["SESS_LAST_ADV_ID"] = intval($_SESSION["SESS_LAST_ADV_ID"] ?? 0)
	return nil
}

func (am AdvModel) FindByByPage(page, cType string) ([]int, string, string, error) {
	strSql := `
		SELECT A.id, A.referer1, A.referer2
		FROM adv A
		INNER JOIN adv_page AP ON (AP.adv_id = A.id and AP.c_type='?')
		WHERE length(AP.page) > 0
		and ? like concat("'%'", AP.page, "'%'")`
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

	// lookup campaign with referer1 and referer2
	$referer1 = trim($referer1)
	$referer1_sql = $referer1 < > '' ? "REFERER1='". $DB- > ForSql($referer1, 255)."'" : "(REFERER1 is null or ". $DB- > Length("REFERER1").
	"=0)"
	$referer2 = trim($referer2)
	$referer2_sql = $referer2 < > '' ? "REFERER2='". $DB- > ForSql($referer2, 255)."'" : "(REFERER2 is null or ". $DB- > Length("REFERER2").
	"=0)"

	sql := `
	SELECT 	ID, REFERER1, REFERER2
	FROM adv
	WHERE  REFERER1=? and REFERER2=?`

	$found = false
	while($wr = $w- > Fetch()) {
		$found = true
		// return with parameters
		$arrADV[] = intval($wr["ID"])
		$ref1 = $wr["REFERER1"]
		$ref2 = $wr["REFERER2"]
	}

	if !$found) {
		$NA = ""
		if COption::GetOptionString("statistic", "ADV_NA") == "Y") {
			$NA_1 = COption::GetOptionString("statistic", "AVD_NA_REFERER1")
			$NA_2 = COption::GetOptionString("statistic", "AVD_NA_REFERER2")
			if ($NA_1 < > '' || $NA_2 < > '') && $referer1 == $NA_1 && $referer2 == $NA_2)
			$NA = "Y"
		}

		if COption::GetOptionString("statistic", "ADV_AUTO_CREATE") == "Y") || ($NA == "Y")) {
if (COption::GetOptionString("statistic", "REFERER_CHECK") == "Y") {
$bGoodR = preg_match("/^([0-9A-Za-z_:;.,-])*$/", $referer1);
if ($bGoodR)
$bGoodR = preg_match("/^([0-9A-Za-z_:;.,-])*$/", $referer2);
} else {
$bGoodR = true;
}

if ($bGoodR) {
// add new advertising campaign
$arFields = array(
"REFERER1" = > $referer1 <> '' ? "'".$DB->ForSql($referer1, 255)."'": "null",
"REFERER2" = > $referer2 <> '' ? "'".$DB->ForSql($referer2, 255)."'": "null",
"DATE_FIRST" = > $DB->GetNowFunction(),
"DATE_LAST" = > $DB->GetNowFunction(),
);
$arrADV[] = $DB->Insert("b_stat_adv", $arFields, $err_mess.__LINE__);
$ref1 = $referer1;
$ref2 = $referer2;
}
}
}
}
