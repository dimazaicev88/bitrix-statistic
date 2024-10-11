package entityjson

type AdvCompany struct {
	AdvNa          bool   `json:"advNa"`          //- Использовать рекламную кампанию c referer1=NA, referer2=NA по умолчанию?
	AdvAutoCreate  bool   `json:"advAutoCreate"`  //- Автоматически создавать рекламные кампании при наличии параметров referer1, referer2 в URL'е:
	RefererCheck   bool   `json:"refererCheck"`   //- Для автоматически создаваемых рекламных кампаний проверять, что referer1 и referer2 содержат только символы латинского алфавита, цифры и символы "_:;.,-":
	SearcherEvents bool   `json:"searcherEvents"` //- Учитывать события рекламных кампаний для поисковиков:
	Referer1Syn    string `json:"referer1Syn"`    //- Синонимы referer1:
	Referer2Syn    string `json:"referer2Syn"`    //- Синонимы referer2:
	Referer3Syn    string `json:"referer3Syn"`    //- Синонимы referer3:
}

type SettingsData struct {
	SaveVisits          bool   `json:"saveVisits"`          //Сохранять статистику посещений разделов и файлов?
	SavePathData        bool   `json:"savePathData"`        //Собирать данные для отчета "Пути по сайту"?
	MaxPathSteps        uint32 `json:"maxPathSteps"`        //Максимальная длина путей по сайту для хранения:
	ImportantPageParams string `json:"importantPageParams"` //Важные параметры страницы:
	DirectoryIndex      string `json:"directoryIndex"`      //Индексные страницы разделов:
	Browsers            string `json:"browsers"`            //Маски для UserAgent браузеров: (используется для автодетекта поисковиков в качестве исключений)
	SaveReferrers       bool   `json:"saveReferrers"`       //Сохранять статистику ссылающихся сайтов и поисковых фраз?
	RefererTop          uint32 `json:"refererTop"`          //Не очищать из ссылающихся доменов первые (TOP):
	SaveHits            bool   `json:"saveHits"`            //Сохранять хиты?
}

type Options struct {
	AdvCompany   *AdvCompany   `json:"advCompany"`
	SettingsData *SettingsData `json:"settingsData"`
}
