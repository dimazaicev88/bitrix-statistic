package entityjson

type LimitActivity struct {
	DefenceDelay        uint64 `json:"defenceDelay"`        //На время
	DefenceStackTime    uint64 `json:"defenceStackTime"`    //Если в течение
	DefenceMaxStackHits uint64 `json:"defenceMaxStackHits"` //Сделано более хитов
	DefenceLog          bool   `json:"defenceLog"`          //Сделать запись в журнал событий
	DefenceOn           bool   `json:"defenceOn"`           //Блокировать
}

type AdvCompany struct {
	AdvNa          bool   `json:"advNa"`          //- Использовать рекламную кампанию c referer1=NA, referer2=NA по умолчанию?
	AdvAutoCreate  bool   `json:"advAutoCreate"`  //- Автоматически создавать рекламные кампании при наличии параметров referer1, referer2 в URL'е:
	RefererCheck   bool   `json:"refererCheck"`   //- Для автоматически создаваемых рекламных кампаний проверять, что referer1 и referer2 содержат только символы латинского алфавита, цифры и символы "_:;.,-":
	SearcherEvents bool   `json:"searcherEvents"` //- Учитывать события рекламных кампаний для поисковиков:
	Referer1Syn    string `json:"referer1Syn"`    //- Синонимы referer1:
	Referer2Syn    string `json:"referer2Syn"`    //- Синонимы referer2:
	Referer3Syn    string `json:"referer3Syn"`    //- Синонимы referer3:
}

type Options struct {
	LimitActivity *LimitActivity `json:"limitActivity"`
	AdvCompany    *AdvCompany    `json:"advCompany"`
}
