package models

import (
	"bitrix-statistic/internal/session"
	"bitrix-statistic/internal/storage"
)

type AdvModel struct {
	storage storage.Storage
}

func (am AdvModel) SetAdv(fullRequestUri string, phpSession *session.Session) {
	//$err_mess = "File: " . __FILE__ . "<br>Line: ";
	//stat_session_register("SESS_ADV_ID") // ID рекламной кампании
	//$DB = CDatabase::GetModuleConnection('statistic');

	// если это начало сессии
	if phpSession.KeyExists("SESS_SESSION_ID") == false || phpSession.GetAsInt("SESS_SESSION_ID") <= 0 &&
		phpSession.KeyExists("SESS_ADV_ID") == false || phpSession.GetAsInt("SESS_ADV_ID") <= 0 {
		var arrADV []string // массив рекламных кампаний

		// проверяем страницу на которую пришел посетитель
		//$page_to = __GetFullRequestUri()
		page, ref1, ref2, err := am.FindByByPage(fullRequestUri, "TO")
		if err != nil {
			return
		}

		// если посетитель пришел с ссылающегося сайта то
		if __GetReferringSite($PROT, $SN, $SN_WithoutPort, $PAGE_FROM)) {
		$site_name = $PROT.$SN;
		// проверяем поисковики
		$strSql = "
		SELECT
		A.REFERER1,
		A.REFERER2,
		S.ADV_ID
		FROM
		b_stat_adv A,
		b_stat_adv_searcher S,
		b_stat_searcher_params P
		WHERE
		S.ADV_ID = A.ID
		and P.SEARCHER_ID = S.SEARCHER_ID
		and upper('" . $DB->ForSql(trim($site_name), 2000) . "')
		like " . $DB->Concat("'%'", "upper(P.DOMAIN)", "'%'") . "
		";
		$w = $DB->Query($strSql, false, $err_mess.__LINE__);
		while ($wr = $w->Fetch()) {
		$ref1 = $wr["REFERER1"];
		$ref2 = $wr["REFERER2"];
		$arrADV[] = intval($wr["ADV_ID"]);
		}

		// проверяем ссылающиеся страницы
		$site_name = $PROT.$SN. $PAGE_FROM;
		CAdv::SetByPage($site_name, $arrADV, $ref1, $ref2, "FROM");
		}

		// если гость пришел с referer1, либо referer2 то
		if ($_SESSION["referer1"] < > '' || $_SESSION["referer2"] < > '') {
		CAdv::SetByReferer(trim($_SESSION["referer1"]), trim($_SESSION["referer2"]), $arrADV, $ref1, $ref2)
		}
		//Handle Openstat if enabled
		if COption::GetOptionString("statistic", "OPENSTAT_ACTIVE") == = "Y" && $_REQUEST["_openstat"] < > '') {
			$openstat = $_REQUEST["_openstat"]
			if mb_strpos($openstat, ";") == = false)
			$openstat = base64_decode($openstat)
			$openstat = explode(";", $openstat)
		CAdv::SetByReferer(
				trim(str_replace(
					array("#service-name#", "#campaign-id#", "#ad-id#", "#source-id#"),
			$openstat,
				COption::GetOptionString("statistic", "OPENSTAT_R1_TEMPLATE")
			)),
			trim(str_replace(
				array("#service-name#", "#campaign-id#", "#ad-id#", "#source-id#"),
			$openstat,
				COption::GetOptionString("statistic", "OPENSTAT_R2_TEMPLATE")
			)),
			$arrADV, $ref1, $ref2
			)
		}
		$arrADV = array_unique($arrADV)

		// если было выявлено более одной рекламной кампании подходящей под условия то
		if count($arrADV) > 1) {
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

	var arrIdAdv []int
	for rows.Next() {
		var id int
		var referer1 string
		var referer2 string
		err = rows.Scan(&id, &referer1, &referer2)
		if err != nil {
			return nil, "", "", err
		}
		arrIdAdv = append(arrIdAdv, id)
	}
	err = rows.Err()
	if err != nil {
		return nil, "", "", err
	}

	return arrIdAdv, page, cType, nil
}
