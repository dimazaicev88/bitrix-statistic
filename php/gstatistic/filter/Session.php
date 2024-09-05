<?php

class Session extends BaseFilter
{

ID* - ID сессии;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
ID_EXACT_MATCH - если значение равно "N", то при фильтрации по ID будет искаться вхождение;
GUEST_ID* - ID посетителя;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
GUEST_ID_EXACT_MATCH - если значение равно "N", то при фильтрации по GUEST_ID будет искаться вхождение;
NEW_GUEST - флаг "новый посетитель", возможные значения:
Y - посетитель впервые на портале;
N - посетитель уже посещал ранее портал.
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
USER_ID* - ID пользователя под которым последний раз был авторизован посетитель;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
USER_ID_EXACT_MATCH - если значение равно "N", то при фильтрации по USER_ID будет искаться вхождение;
USER_AUTH - флаг "был ли посетитель авторизован в данной сессии", возможные значения:
Y - да;
N - нет.
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
USER* - ID, логин, имя, фамилия пользователя под которым последний раз был авторизован посетитель;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
USER_EXACT_MATCH - если значение равно "Y", то при фильтрации по USER будет искаться точное совпадение;
REGISTERED - флаг "был ли авторизован посетитель в данной сессии или до этого", возможные значения:
Y - был;
N - не был.
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
FAVORITES - флаг "добавлял ли посетитель сайт в "Избранное"", возможные значения:
Y - да;
N - нет.
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
EVENTS1 - начальное значение интервала для поля "количество событий данной сессии";
EVENTS2 - конечное значение интервала для поля "количество событий данной сессии";
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
HITS1 - начальное значение интервала для поля "количество хитов данной сессии";
HITS2 - конечное значение интервала для поля "количество хитов данной сессии";
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
ADV - флаг "приходил ли посетитель в данной сессии по какой-либо рекламной кампании", возможные значения:
Y - да;
N - нет.
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
ADV_ID* - ID рекламной кампании;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
ADV_ID_EXACT_MATCH - если значение равно "N", то при фильтрации по ADV_ID будет искаться вхождение;
ADV_BACK - флаг "возврат по рекламной кампании", возможные значения:
Y - был возврат;
N - был прямой заход.
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
REFERER1* - идентификатор referer1 рекламной кампании;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
REFERER1_EXACT_MATCH - если значение равно "Y", то при фильтрации по REFERER1 будет искаться точное совпадение;
REFERER2* - идентификатор referer2 рекламной кампании;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
REFERER2_EXACT_MATCH - если значение равно "Y", то при фильтрации по REFERER2 будет искаться точное совпадение;
REFERER3* - дополнительный параметр referer3 рекламной кампании;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
REFERER3_EXACT_MATCH - если значение равно "Y", то при фильтрации по REFERER3 будет искаться точное совпадение;
STOP - флаг "попал ли посетитель под какую либо запись стоп-листа", возможные значения:
Y - да;
N - нет.
STOP_LIST_ID* - ID записи стоп-листа под которую попал посетитель, если это имело место быть;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
STOP_LIST_ID_EXACT_MATCH - если значение равно "N", то при фильтрации по STOP_LIST_ID будет искаться вхождение;
COUNTRY_ID* - ID страны посетителя;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
COUNTRY_ID_EXACT_MATCH - если значение равно "N", то при фильтрации по COUNTRY_ID будет искаться вхождение;
COUNTRY* - наименование страны посетителя;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
COUNTRY_EXACT_MATCH - если значение равно "Y", то при фильтрации по COUNTRY будет искаться точное совпадение;
IP* - IP адрес посетителя на последнем хите сессии;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
IP_EXACT_MATCH - если значение равно "Y", то при фильтрации по IP будет искаться точное совпадение;
USER_AGENT* - UserAgent посетителя;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
USER_AGENT_EXACT_MATCH - если значение равно "Y", то при фильтрации по USER_AGENT будет искаться точное совпадение;
DATE_START_1 - начальное значение интервала для поля "время первого хита сессии";
    /**
     * Значение интервала для поля "время первого хита сессии"
     *
     * @param Operator $operator
     * @param string $date
     * @return $this
     */
    public function date(Operator $operator, string $date): PhraseFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
DATE_END_1 - начальное значение интервала для поля "время последнего хита сессии";
DATE_END_2 - конечное значение интервала для поля "время последнего хита сессии";
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
URL_TO* - первая страница сессии;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
URL_TO_EXACT_MATCH - если значение равно "Y", то при фильтрации по URL_TO будет искаться точное совпадение;
URL_TO_404 - была ли 404 ошибка на первой страницы сессии
Y - была;
N - не было.
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
FIRST_SITE_ID* - ID сайта на первом хите сессии;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
FIRST_SITE_ID_EXACT_MATCH - если значение равно "N", то при фильтрации по FIRST_SITE_ID будет искаться вхождение;
URL_LAST* - последняя страница сессии;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
URL_LAST_EXACT_MATCH - если значение равно "Y", то при фильтрации по URL_LAST будет искаться точное совпадение;
URL_LAST_404 - была ли 404 ошибка на последней страницы сессии
Y - была;
N - не было.
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
LAST_SITE_ID* - ID сайта на последнем хите сессии;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
}