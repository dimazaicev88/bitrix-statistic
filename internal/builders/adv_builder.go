package builders

//var selectBuffer = map[string]string{
//	"ID":                           "", //- ID РК;
//	"PRIORITY":                     "", // - приоритет;
//	"REFERER1":                     "", // - идентификатор referer1 РК;
//	"REFERER2":                     "", // - идентификатор referer2 РК;
//	"C_TIME_FIRST":                 "", // - время начала РК (первый прямой заход);
//	"C_TIME_LAST":                  "", // - последний прямой заход или возврат по РК;
//	"ADV_TIME":                     "", // - длительность РК (разница C_TIME_LAST - C_TIME_FIRST);
//	"ATTENT":                       "", // - коэфициент внимательности посетителей на прямом заходе по РК;
//	"ATTENT_BACK":                  "", // - коэфициент внимательности посетителей на возврате по РК;
//	"NEW_VISITORS":                 "", // - процент посетителей впервые пришедших на сайт по данной РК от общего количества посетителей пришедших по данной РК;
//	"RETURNED_VISITORS ":           "", //- процент посетителей возвратившихся на сайт после прямого захода по данной РК;
//	"VISITORS_PER_DAY":             "", // - среднее количество посетителей за день;
//	"COST":                         "", // - затраты на РК;
//	"REVENUE ":                     "", //- доходы с РК;
//	"BENEFIT ":                     "", //- прибыль РК;
//	"ROI":                          "", // - рентабельность РК;
//	"SESSION_COST":                 "", // - средняя стоимость сессии (затраты/кол-во сессий на прямом заходе);
//	"VISITOR_COST ":                "", //- средняя стоимость посетителя (затраты/кол-во посетителей на прямых заходах);
//	"GUESTS":                       "", // - суммарное кол-во посетителей на прямых заходах;
//	"GUESTS_BACK":                  "", // - суммарное кол-во посетителей на возвратах;
//	"NEW_GUESTS":                   "", // - суммарное кол-во новых посетителей по данной РК;
//	"FAVORITES":                    "", // - суммарное кол-во посетителей, добавивших сайт в "Избранное" на прямом заходе по РК;
//	"FAVORITES_BACK":               "", // - суммарное кол-во посетителей, добавивших сайт в "Избранное" на возврате по РК;
//	"C_HOSTS":                      "", // - суммарное кол-во хостов на прямом заходе по РК;
//	"HOSTS_BACK ":                  "", //- суммарное кол-во хостов на возврате по РК;
//	"SESSIONS":                     "", // - суммарное кол-во сессий на прямом заходе по РК;
//	"SESSIONS_BACK ":               "", //- суммарное кол-во сессий на возврате по РК;
//	"HITS":                         "", // - суммарное кол-во хитов на прямом заходе по РК;
//	"HITS_BACK":                    "", // - суммарное кол-во хитов на возврате по РК;
//	"GUESTS_TODAY ":                "", //- кол-во посетителей на прямом заходе за сегодня;
//	"GUESTS_BACK_TODAY":            "", // - кол-во посетителей на возврате за сегодня;
//	"NEW_GUESTS_TODAY":             "", // - кол-во новых посетителей за сегодня;
//	"FAVORITES_TODAY":              "", // - кол-во посетителей, добавивших сайт в "Избранное" на прямом заходе за сегодня;
//	"FAVORITES_BACK_TODAY":         "", // - кол-во посетителей, добавивших сайт в "Избранное" на возврате за сегодня;
//	"C_HOSTS_TODAY":                "", // - кол-во хостов на прямом заходе за сегодня;
//	"HOSTS_BACK_TODAY":             "", // - кол-во хостов на возврате за сегодня;
//	"SESSIONS_TODAY":               "", // - кол-во сессий на прямом заходе за сегодня;
//	"SESSIONS_BACK_TODAY":          "", // - кол-во сессий на возврате за сегодня;
//	"HITS_TODAY":                   "", // - кол-во хитов на прямом заходе за сегодня;
//	"HITS_BACK_TODAY ":             "", //- кол-во хитов на возврате за сегодня;
//	"GUESTS_YESTERDAY ":            "", //- кол-во посетителей на прямом заходе за вчера;
//	"GUESTS_BACK_YESTERDAY":        "", // - кол-во посетителей на возврате за вчера;
//	"NEW_GUESTS_YESTERDAY":         "", // - кол-во новых посетителей за вчера;
//	"FAVORITES_YESTERDAY":          "", // - кол-во посетителей, добавивших сайт в "Избранное" на прямом заходе за вчера;
//	"FAVORITES_BACK_YESTERDAY":     "", // - кол-во посетителей, добавивших сайт в "Избранное" на возврате за вчера;
//	"C_HOSTS_YESTERDAY":            "", // - кол-во хостов на прямом заходе за вчера;
//	"HOSTS_BACK_YESTERDAY":         "", // - кол-во хостов на возврате за вчера;
//	"SESSIONS_YESTERDAY":           "", // - кол-во сессий на прямом заходе за вчера;
//	"SESSIONS_BACK_YESTERDAY ":     "", //- кол-во сессий на возврате за вчера;
//	"HITS_YESTERDAY":               "", // - кол-во хитов на прямом заходе за вчера;
//	"HITS_BACK_YESTERDAY":          "", // - кол-во хитов на возврате за вчера;
//	"GUESTS_BEF_YESTERDAY":         "", // - кол-во посетителей на прямом заходе за позавчера;
//	"GUESTS_BACK_BEF_YESTERDAY":    "", // - кол-во посетителей на возврате за позавчера;
//	"NEW_GUESTS_BEF_YESTERDAY":     "", // - кол-во новых посетителей за позавчера;
//	"FAVORITES_BEF_YESTERDAY":      "", // - кол-во посетителей, добавивших сайт в "Избранное" на прямом заходе за позавчера;
//	"FAVORITES_BACK_BEF_YESTERDAY": "", // - кол-во посетителей, добавивших сайт в "Избранное" на возврате за позавчера;
//	"C_HOSTS_BEF_YESTERDAY":        "", // - кол-во хостов на прямом заходе за позавчера;
//	"HOSTS_BACK_BEF_YESTERDAY":     "", // - кол-во хостов на возврате за позавчера;
//	"SESSIONS_BEF_YESTERDAY":       "", // - кол-во сессий на прямом заходе за позавчера;
//	"SESSIONS_BACK_BEF_YESTERDAY":  "", // - кол-во сессий на возврате за позавчера;
//	"HITS_BEF_YESTERDAY":           "", // - кол-во хитов на прямом заходе за позавчера;
//	"HITS_BACK_BEF_YESTERDAY":      "", // - кол-во хитов на возврате за позавчера;
//	"GUESTS_PERIOD ":               "", //- кол-во посетителей на прямом заходе за установленный в фильтре (filter) интервал времени;
//	"GUESTS_BACK_PERIOD":           "", // - кол-во посетителей на возврате за установленный в фильтре интервал времени;
//	"NEW_GUESTS_PERIOD ":           "", //- кол-во новых посетителей за установленный в фильтре интервал времени;
//	"FAVORITES_PERIOD":             "", // - кол-во посетителей, добавивших сайт в "Избранное" на прямом заходе за установленный в фильтре интервал времени;
//	"FAVORITES_BACK_PERIOD ":       "", //- кол-во посетителей, добавивших сайт в "Избранное" на возврате за установленный в фильтре интервал времени;
//	"C_HOSTS_PERIOD ":              "", //- кол-во хостов на прямом заходе за установленный в фильтре интервал времени;
//	"HOSTS_BACK_PERIOD":            "", // - кол-во хостов на возврате за установленный в фильтре интервал времени;
//	"SESSIONS_PERIOD":              "", // - кол-во сессий на прямом заходе за установленный в фильтре интервал времени;
//	"SESSIONS_BACK_PERIOD ":        "", //- кол-во сессий на возврате за установленный в фильтре интервал времени;
//	"HITS_PERIOD":                  "", // - кол-во хитов на прямом заходе за установленный в фильтре интервал времени;
//	"HITS_BACK_PERIOD":             "", // - кол-во хитов на возврате за установленный в фильтре интервал времени.,
//}
//
//var filterFields = map[string]string{
//	"GROUP": "", // - список возможных значений:
//	//referer1 - список РК будет сгруппирован по идентификатору referer1 РК;
//	//referer2 - список РК будет сгруппирован по идентификатору referer2 РК.
//	"ID":           "", // - ID РК;
//	"DATE1_PERIOD": "", // - начальное значение периода за который необходимо получить данные;
//	"DATE2_PERIOD": "", // - конечное значение периода за который необходимо получить данные;
//	"DATE1_FIRST":  "", // - начальное значение интервала для поля "время начала РК";
//	"DATE2_FIRST":  "", // - конечное значение интервала для поля "время начала РК";
//	"DATE1_LAST":   "", // - начальное значение интервала для поля "время окончания РК";
//	"DATE2_LAST":   "", // конечное значение интервала для поля "время окончания РК";
//	"REFERER1*":    "", // идентификатор referer1 РК;
//	"REFERER2*":    "", // идентификатор referer2 РК;
//	"PRIORITY1":    "", // начальное значение интервала для поля "приоритет РК";
//	"PRIORITY2":    "", // конечное значение интервала для поля "приоритет РК";
//	"NEW_GUESTS1":  "", // начальное значение интервала для поля "новые посетители РК";
//	"NEW_GUESTS2":  "", // конечное значение интервала для поля "новые посетители РК";
//	"GUESTS_BACK":  "", // флаг, означающий по какому полю, фильтровать посетителей, список возможных значений:
//	// N - на прямом заходе;
//	//Y - на возврате.
//	"GUESTS1":        "", //начальное значение интервала для поля "посетители на прямом заходе или на возврате" (в зависимости от флага GUESTS_BACK);
//	"GUESTS2":        "", //конечное значение интервала для поля "посетители на прямом заходе или на возврате" (в зависимости от флага GUESTS_BACK);
//	"FAVORITES_BACK": "", //флаг, означающий по какому полю, фильтровать посетителей, добавившие сайт в "Избранное", список возможных значений:
//	//N - на прямом заходе;
//	//Y - на возврате.
//	"FAVORITES1": "", //начальное значение интервала для поля "посетители, добавившие сайт в Избранное на прямом заходе или возврате" (в зависимости от флага FAVORITES_BACK);
//	"FAVORITES2": "", //конечное значение интервала для поля "посетители, добавившие сайт в Избранное на прямом заходе или возврате" (в зависимости от флага FAVORITES_BACK);
//	"HOSTS_BACK": "", //флаг, означающий по какому полю, фильтровать хосты, список возможных значений:
//	//N - на прямом заходе;
//	//Y - на возврате.
//	"HOSTS1":        "", //начальное значение интервала для поля "хосты на прямом заходе или возврате" (в зависимости от флага HOSTS_BACK);
//	"HOSTS2":        "", //конечное значение интервала для поля "хосты на прямом заходе или возврате" (в зависимости от флага HOSTS_BACK);
//	"SESSIONS_BACK": "", //флаг, означающий по какому полю, фильтровать сессии, список возможных значений:
//	//N - на прямом заходе;
//	//Y - на возврате.
//	"SESSIONS1": "", //начальное значение интервала для поля "сессии на прямом заходе или возврате" (в зависимости от флага SESSIONS_BACK);
//	"SESSIONS2": "", //конечное значение интервала для поля "сессии на прямом заходе или возврате" (в зависимости от флага SESSIONS_BACK);
//	"HITS_BACK": "", //флаг, означающий по какому полю, фильтровать хиты, список возможных значений:
//	//N - на прямом заходе;
//	//Y - на возврате.
//	"HITS1":             "", //начальное значение интервала для поля "хиты на прямом заходе или возврате" (в зависимости от флага HITS_BACK);
//	"HITS2":             "", //конечное значение интервала для поля "хиты на прямом заходе или возврате" (в зависимости от флага HITS_BACK);
//	"COST1":             "", //начальное значение интервала для поля "затраты на РК";
//	"COST2":             "", //конечное значение интервала для поля "затраты на РК";
//	"REVENUE1":          "", //начальное значение интервала для поля "доходы с РК";
//	"REVENUE2":          "", //конечное значение интервала для поля "доходы с РК";
//	"BENEFIT1":          "", //начальное значение интервала для поля "прибыль с РК";
//	"BENEFIT2":          "", //конечное значение интервала для поля "прибыль с РК";
//	"ROI1":              "", //начальное значение интервала для поля "рентабельность РК";
//	"ROI2":              "", //конечное значение интервала для поля "рентабельность РК";
//	"ATTENT1":           "", //начальное значение интервала для поля "коэффициент внимательности посетителей РК";
//	"ATTENT2":           "", //конечное значение интервала для поля "коэффициент внимательности посетителей РК";
//	"VISITORS_PER_DAY1": "", //начальное значение интервала для поля "среднее кол-во посетителей в день";
//	"VISITORS_PER_DAY2": "", //конечное значение интервала для поля "среднее кол-во посетителей в день";
//	"DURATION1":         "", //начальное значение интервала для поля "длительность РК";
//	"DURATION2":         "", //конечное значение интервала для поля "длительность РК";
//	"CURRENCY":          "", //валюта в которой заданы финансовые показатели РК;
//	"DESCRIPTION*":      "", //описание РК;
//}
