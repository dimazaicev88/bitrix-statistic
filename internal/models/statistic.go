package models

import (
	"bitrix-statistic/internal/entity"
	"bitrix-statistic/internal/session"
	"bitrix-statistic/internal/storage"
	"github.com/sirupsen/logrus"
	"strconv"
)

type StatisticModel struct {
	storage     storage.Storage
	guestModel  *GuestModel
	optionModel *OptionModel
	advModel    *AdvModel
	logger      logrus.Logger
}

func NewStatisticModel(storage storage.Storage) StatisticModel {
	return StatisticModel{
		storage:     storage,
		guestModel:  NewGuestModel(storage),
		optionModel: NewOptionModel(storage),
		advModel:    NewAdvModel(storage, NewOptionModel(storage)),
	}
}

func (stm *StatisticModel) Add(data entity.StatData) error {
	guestDb, err := stm.guestModel.FindByToken(data.Token)
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

func (stm *StatisticModel) SetGuest(phpSession *session.Session, siteId, referrer, fullRequestUrl, error404 string, cookieGuestId, cookieLastVisit, cookieAdvId int) (string, string, error) {
	phpSession.Set("SESS_GUEST_ID", "")        // ID гостя
	phpSession.Set("SESS_GUEST_NEW", "")       // флаг "новый гость"
	phpSession.Set("SESS_LAST_USER_ID", "")    // под кем гость был авторизован в последний раз
	phpSession.Set("SESS_LAST_ADV_ID", "")     // по какой рекламной кампании был в последний раз
	phpSession.Set("SESS_GUEST_FAVORITES", "") // флаг добавлял ли гость сайт в фавориты
	phpSession.Set("SESS_LAST", "")            // Y - гость сегодня уже заходил; N - еще не заходил

	lastReferer1 := ""
	lastReferer2 := ""

	repairCookieGuest := "N"
	if phpSession.KeyExists("SESS_GUEST_NEW") == false {
		phpSession.Set("SESS_GUEST_NEW", "N")
	}
	phpSession.Set("SESS_GUEST_ID", strconv.Itoa(phpSession.GetAsInt("SESS_GUEST_ID")))

	if cookieGuestId == 0 {
		cookieGuestId = phpSession.GetAsInt("SESS_GUEST_ID")
	}
	// если сессия только открылась
	if phpSession.KeyExists("SESS_SESSION_ID") == false || phpSession.GetAsInt("SESS_SESSION_ID") <= 0 {
		// выбираем из базы параметры гостя
		guestId, favorites, lastUserId, lastAdvId, last, err := stm.guestModel.FindLastById(cookieGuestId)
		if err != nil {
			return "", "", err
		}
		// если ничего не выбрали то
		if guestId == 0 {
			// считаем гостя новым
			phpSession.Set("SESS_GUEST_ID", "")
			phpSession.Set("SESS_GUEST_NEW", "Y")
			phpSession.Set("SESS_GUEST_FAVORITES", "N")

			// если у него в cookie хранится GUEST_ID то
			if cookieGuestId > 0 {
				phpSession.Set("SESS_GUEST_NEW", "N")
				// получаем дату последнего посещения сайта данным гостем
				// если формат корректный то
				if cookieLastVisit > 0 {
					// получаем дату последней инсталляции таблиц модуля
					dateInstall := stm.optionModel.GetWithDefault("INSTALL_STATISTIC_TABLES", "NOT_FOUND")
					if dateInstall == "NOT_FOUND" {
						//TODO ?
						//dateInstall = date("d.m.Y H:i:s", time());
						//stm.optionModel.Set("INSTALL_STATISTIC_TABLES", $dateInstall, "Installation date of Statistics module tables");
					}
					//TODO
					//if ($dateInstall = MkDateTime($dateInstall, "d.m.Y H:i:s")) {
					//// если таблицы были инсталлированы после последнего посещения сайта то
					//if ($DATE_INSTALL > $LAST_VISIT) {
					//// Посетитель считается новым т.к. он нигде не был учтен
					//$_SESSION["SESS_GUEST_NEW"] = "Y";
					//}
					//}
				}
				// устанавливаем флаг того что мы восстанавливаем гостя
				//repairCookieGuest := "Y"
				//получаем идентификатор его последней рекламной кампании
				//$CookieAdv = $GLOBALS["APPLICATION"]- > get_cookie("LAST_ADV")
			}
		} else // иначе если выбрали параметры гостя то
		{
			// то запоминаем их в сессии
			phpSession.Set("SESS_GUEST_FAVORITES", favorites)
			//phpSession.Set("SESS_GUEST_FAVORITES") = ($_SESSION["SESS_GUEST_FAVORITES"] == "Y") ? "Y": "N"
			if phpSession.KeyExists("SESS_GUEST_NEW") == false {
				phpSession.Set("SESS_GUEST_NEW", "N")
			}
			phpSession.Set("SESS_GUEST_ID", strconv.Itoa(guestId))
			phpSession.Set("SESS_LAST_ADV_ID", strconv.Itoa(lastAdvId))
			phpSession.Set("SESS_LAST_USER_ID", strconv.Itoa(lastUserId))
			phpSession.Set("SESS_LAST", last)
			if phpSession.GetAsInt("SESS_LAST_ADV_ID") > 0 {
				sql := `SELECT referer1, referer2 FROM adv WHERE id=?`
				rows, err := stm.storage.DB().Query(sql, phpSession.GetAsInt("SESS_LAST_ADV_ID"))
				if err != nil {
					return "", "", nil
				}

				for rows.Next() {
					err = rows.Scan(&lastReferer1, &lastReferer2)
					if err != nil {
						return "", "", nil
					}
				}
				err = rows.Err()
				if err != nil {
					return "", "", nil
				}
			}
		}
	}
	// если есть необходимость то
	if phpSession.GetAsInt("SESS_GUEST_ID") <= 0 {
		guestData := entity.GuestDb{
			//UrlFrom:  referrer,
			//UrlTo:    fullRequestUrl,
			//UrlTo404: error404,
			//SiteId:   siteId,
			//AdvId:    phpSession.GetAsInt("SESS_ADV_ID"),
			//Referer1: phpSession.Get("referer1"),
			//Referer2: phpSession.Get("referer2"),
			//Referer3: phpSession.Get("referer3"),
		}
		// если мы восстанавливаем гостя по данным записанным в его cookie то
		if repairCookieGuest == "Y" {
			// если гость не считается новым, то добавим ему одну сессию
			if phpSession.Get("SESS_GUEST_NEW") == "N" {
				guestData.Sessions = 1
			}
			if cookieAdvId > 0 {
				adv, err := stm.advModel.FindById(cookieAdvId)
				if err != nil {
					return "", "", err
				}
				// если в базе есть такая рекламная кампания то
				if (adv != entity.AdvDb{}) {
					// считаем что гость вернулся по данной рекламной кампании
					phpSession.Set("SESS_LAST_ADV_ID", strconv.Itoa(cookieAdvId))

					// если последний вход записанный в cookie
					// не был прямым входом по рекламной кампании то
					//guestData.FirstAdvId = cookieAdvId
					//guestData.FirstReferer1 = adv.Referer1
					//guestData.FirstReferer2 = adv.Referer2
					//guestData.LastAdvId = cookieAdvId
					//guestData.LastAdvBack = "Y"
					//guestData.LastReferer1 = adv.Referer1
					//guestData.LastReferer2 = adv.Referer2
					lastReferer1 = adv.Referer1
					lastReferer2 = adv.Referer2
				}
			}
		}
		stm.guestModel.Add(guestData)
		//$_SESSION["SESS_GUEST_ID"] = $DB- > Insert("b_stat_guest", $arFields, $err_mess.__LINE__)
		//if ($ERROR_404 == "N") {
		//CStatistics::Set404("b_stat_guest", "ID = ".intval($_SESSION["SESS_GUEST_ID"]), array("FIRST_URL_TO_404" = > "Y"))
		//}
	}

	//TODO вынести в phpCookie
	//// если гость авторизовался то
	//if is_object($USER) && intval($USER- > GetID()) > 0) {
	//	// запоминаем кто он
	//	$_SESSION["SESS_LAST_USER_ID"] = intval($USER- > GetID())
	//}
	//if intval($_SESSION["SESS_LAST_USER_ID"] ?? 0) <= 0) {
	//	$_SESSION["SESS_LAST_USER_ID"] = ""
	//}
	//
	//if ($_SESSION["SESS_GUEST_ID"] > 0) {
	//	// сохраним ID посетителя в куках
	//	$GLOBALS["APPLICATION"]- > set_cookie("GUEST_ID", $_SESSION["SESS_GUEST_ID"])
	//}
	//// сохраним в cookie дату последнего посещения данным гостем сайта
	//$GLOBALS["APPLICATION"]- > set_cookie("LAST_VISIT", date("d.m.Y H:i:s", time()))

	return lastReferer1, lastReferer2, nil
}

func (stm *StatisticModel) SetNewDay() {

}
