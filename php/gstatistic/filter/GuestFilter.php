<?php

class GuestFilter extends BaseFilter
{

    /**
     * @param Operator $operator
     * @param int $value
     * @return $this
     */
    function id(Operator $operator, int $value): GuestFilter
    {
        $this->setFilter($operator, $value, "id");
        return $this;
    }

    /**
     * Был ли посетитель когда-либо авторизован на сайте
     * @param bool $value
     * @return $this
     */
    function isRegistered(bool $value): GuestFilter
    {
        $this->setFilter(Operator::Eq, intval($value), "registered");
        return $this;
    }

    /**
     * Начальное значение интервала для поля "дата первого захода на сайт"
     * @param Operator $operator
     * @param string $value
     * @return GuestFilter
     */
    function firstDate1(Operator $operator, string $value): GuestFilter
    {
        $this->setFilter($operator, $value, "first_date_1");
        return $this;
    }

    /**
     * Конечное значение интервала для поля "дата первого захода на сайт"
     * @param Operator $operator
     * @param string $value
     * @return GuestFilter
     */
    function firstDate2(Operator $operator, string $value): GuestFilter
    {
        $this->setFilter($operator, $value, "first_date_2");
        return $this;
    }

    /**
     * Начальное значение интервала для поля "дата последнего захода на сайт"
     * @param Operator $operator
     * @param string $value
     * @return GuestFilter
     */
    function lastDate1(Operator $operator, string $value): GuestFilter
    {
        $this->setFilter($operator, $value, "last_date_1");
        return $this;
    }

    /**
     * Конечное значение интервала для поля "дата первого захода на сайт"
     * @param Operator $operator
     * @param string $value
     * @return GuestFilter
     */
    function lastDate2(Operator $operator, string $value): GuestFilter
    {
        $this->setFilter($operator, $value, "last_date_2");
        return $this;
    }

    /**
     * Начальное значение интервала для даты посещения посетителем сайта
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    function periodDate1(Operator $operator, string $value): GuestFilter
    {
        $this->setFilter($operator, $value, "period_date_1");
        return $this;
    }

    /**
     * Конечно значение интервала для даты посещения посетителем сайта
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    function periodDate2(Operator $operator, string $value): GuestFilter
    {
        $this->setFilter($operator, $value, "period_date_2");
        return $this;
    }

    /**
     * ID сайта первого либо последнего захода;
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    function siteId(Operator $operator, string $value): GuestFilter
    {
        $this->setFilter($operator, $value, "site_id");
        return $this;
    }

    /**
     * ID сайта первого захода
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    function firstSiteId(Operator $operator, string $value): GuestFilter
    {
        $this->setFilter($operator, $value, "first_site_id");
        return $this;
    }


    /**
     * ID сайта последнего захода
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    function lastSiteId(Operator $operator, string $value): GuestFilter
    {
        $this->setFilter($operator, $value, "last_site_id");
        return $this;
    }

    /**
     * Страница откуда впервые пришел посетитель, страница на которую впервые пришел посетитель и последняя страница просмотренная посетителем;
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    function url(Operator $operator, string $value): GuestFilter
    {
        $this->setFilter($operator, $value, "url");
        return $this;
    }

    /**
     * URL_404 - была ли 404 ошибка на первой странице или на последней странице посещенной посетителем:
     * @return $this
     */
    function existsUrl404(bool $value): GuestFilter
    {
        $this->setFilter(Operator::Eq, intval($value), "url_404");
        return $this;
    }

    /**
     * UserAgent посетителя на последнем заходе;
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    function userAgent(Operator $operator, string $value): GuestFilter
    {
        $this->setFilter($operator, $value, "user_agent");
        return $this;
    }

    /**
     * Флаг "приходил ли посетитель когда-либо по рекламной кампании (не равной NA/NA)":
     * true - посетитель приходил по какой-либо рекламной кампании(не равной NA / NA);
     * false - не приходил никогда ни по одной рекламной кампании(не равной NA / NA).
     * @return $this
     */
    function isAdv(bool $value): GuestFilter
    {
        $this->setFilter(Operator::Eq, intval($value), "adv");
        return $this;
    }

    /**
     * ID рекламной кампании первого либо последнего захода посетителя(при этом это мог быть как прямой заход, так и возврат по рекламной кампании)
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    function advUuid(Operator $operator, string $value): GuestFilter
    {
        $this->setFilter($operator, $value, "advUuid");
        return $this;
    }

    /**
     * Идентификатор referer1 рекламной кампании первого либо последнего захода посетителя;
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    function referer1(Operator $operator, string $value): GuestFilter
    {
        $this->setFilter($operator, $value, "referer1");
        return $this;
    }

    /**
     * Идентификатор referer2 рекламной кампании первого либо последнего захода посетителя
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    function referer2(Operator $operator, string $value): GuestFilter
    {
        $this->setFilter($operator, $value, "referer2");
        return $this;
    }

    /**
     * Дополнительный параметр referer3 рекламной кампании первого либо последнего захода посетителя
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    function referer3(Operator $operator, string $value): GuestFilter
    {
        $this->setFilter($operator, $value, "referer3");
        return $this;
    }

    /**
     * Количество событий сгенерированных посетителем
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    function events(Operator $operator, string $value): GuestFilter
    {
        $this->setFilter($operator, $value, "events");
        return $this;
    }


    /**
     * Количество сессий сгенерированных посетителем
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    function sessions(Operator $operator, string $value): GuestFilter
    {
        $this->setFilter($operator, $value, "sessions");
        return $this;
    }

    /**
     * Количество хитов сгенерированных посетителем;
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    function hits(Operator $operator, string $value): GuestFilter
    {
        $this->setFilter($operator, $value, "hits_1");
        return $this;
    }


    /**
     * Флаг "добавлял ли посетитель сайт в "Избранное"":
     * true - добавлял;
     * false - не добавлял
     * @return $this
     */
    function isAddToFavorites(bool $value): GuestFilter
    {
        $this->setFilter(Operator::Eq, intval($value), "favorite");
        return $this;
    }

    /**
     * IP адрес посетителя сайта в последнем заходе
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    function ip(Operator $operator, string $value): GuestFilter
    {
        $this->setFilter($operator, intval($value), "ip");
        return $this;
    }

    /**
     * Языки установленные в настройках браузера посетителя в последнем заходе
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    function lang(Operator $operator, string $value): GuestFilter
    {
        $this->setFilter($operator, $value, "lang");
        return $this;
    }

    /**
     * ID страны(двух символьный идентификатор) посетителя в последнем заходе
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    function countryId(Operator $operator, string $value): GuestFilter
    {
        $this->setFilter($operator, $value, "countryId");
        return $this;
    }

    /**
     * Название страны
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    function country(Operator $operator, string $value): GuestFilter
    {
        $this->setFilter($operator, $value, "country");
        return $this;
    }

    /**
     * ID, логин, имя, фамилия пользователя, под которыми посетитель последний раз был авторизован
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    function user(Operator $operator, string $value): GuestFilter
    {
        $this->setFilter($operator, $value, "user");
        return $this;
    }

    /**
     * ID пользователя, под которым посетитель последний раз был авторизован
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    function userId(Operator $operator, string $value): GuestFilter
    {
        $this->setFilter($operator, $value, "userId");
        return $this;
    }
}