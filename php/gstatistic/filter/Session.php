<?php

class Session extends BaseFilter
{

    /**
     * UUID сессии;
     *
     * @param Operator $operator
     * @param string $value
     * @return Session
     */
    public function uuid(Operator $operator, string $value): Session
    {
        $this->setFilter($operator, $value, 'uuid');
        return $this;
    }

    /**
     * UUID посетителя;
     *
     * @param Operator $operator
     * @param string $value
     * @return Session
     */
    public function guestUuid(Operator $operator, string $value): Session
    {
        $this->setFilter($operator, $value, 'guestUuid');
        return $this;
    }

    /**
     * Флаг "новый посетитель"
     *
     * @param bool $value
     * @return Session
     */
    public function isNewGuest(bool $value): Session
    {
        $this->setFilter(Operator::Eq, $value, 'isNewGuest');
        return $this;
    }

    /**
     * ID пользователя под которым последний раз был авторизован посетитель;
     *
     * @param Operator $operator
     * @param string $value
     * @return Session
     */
    public function userId(Operator $operator, string $value): Session
    {
        $this->setFilter($operator, $value, 'userId');
        return $this;
    }

    /**
     * Флаг "был ли посетитель авторизован в данной сессии"
     *
     * @param Operator $operator
     * @param bool $value
     * @return Session
     */
    public function isUserAuth(Operator $operator, bool $value): Session
    {
        $this->setFilter(Operator::Eq, $value, 'isUserAuth');
        return $this;
    }

    /**
     * Флаг "был ли авторизован посетитель в данной сессии или до этого"
     *
     *
     * @param bool $value
     * @return Session
     */
    public function isRegistered(bool $value): Session
    {
        $this->setFilter(Operator::Eq, $value, 'isRegistered');
        return $this;
    }

    /**
     * Флаг "добавлял ли посетитель сайт в "Избранное""
     *
     * @param bool $value
     * @return Session
     */
    public function isFavorites(bool $value): Session
    {
        $this->setFilter(Operator::Eq, $value, 'isFavorites');
        return $this;
    }

    /**
     * Количество событий данной сессии
     *
     * @param Operator $operator
     * @param int $value
     * @return Session
     */
    public function events(Operator $operator, int $value): Session
    {
        $this->setFilter($operator, $value, 'events');
        return $this;
    }


    /**
     * Количество хитов данной сессии
     *
     * @param Operator $operator
     * @param int $value
     * @return Session
     */
    public function hits(Operator $operator, int $value): Session
    {
        $this->setFilter($operator, $value, 'hits');
        return $this;
    }

    /**
     * Приходил ли посетитель в данной сессии по какой-либо рекламной кампании
     *
     * @param bool $value
     * @return Session
     */
    public function isAdv(bool $value): Session
    {
        $this->setFilter(Operator::Eq, $value, 'isAdv');
        return $this;
    }

    /**
     * UUID рекламной кампании;
     *
     * @param Operator $operator
     * @param string $value
     * @return Session
     */
    public function advUuid(Operator $operator, string $value): Session
    {
        $this->setFilter($operator, $value, 'advUuid');
        return $this;
    }

    /**
     * Флаг "возврат по рекламной кампании"
     *
     * @param Operator $operator
     * @param bool $value
     * @return Session
     */
    public function isAdvBack(Operator $operator, bool $value): Session
    {
        $this->setFilter($operator, $value, 'isAdvBack');
        return $this;
    }

    /**
     * Идентификатор referer1 рекламной кампании
     *
     * @param Operator $operator
     * @param string $value
     * @return Session
     */
    public function referer1(Operator $operator, string $value): Session
    {
        $this->setFilter($operator, $value, 'referer1');
        return $this;
    }

    /**
     * Идентификатор referer2 рекламной кампании
     *
     * @param Operator $operator
     * @param string $value
     * @return Session
     */
    public function referer2(Operator $operator, string $value): Session
    {
        $this->setFilter($operator, $value, 'referer2');
        return $this;
    }

    /**
     * Дополнительный параметр referer3 рекламной кампании;
     *
     * @param Operator $operator
     * @param string $value
     * @return Session
     */
    public function referer3(Operator $operator, string $value): Session
    {
        $this->setFilter($operator, $value, 'referer3');
        return $this;
    }

    /**
     * Флаг "попал ли посетитель под какую либо запись стоп-листа"
     *
     * @param Operator $operator
     * @param string $value
     * @return Session
     */
    public function stop(Operator $operator, string $value): Session
    {
        $this->setFilter($operator, $value, 'stop');
        return $this;
    }

    /**
     * ID записи стоп-листа под которую попал посетитель
     *
     * @param Operator $operator
     * @param string $value
     * @return Session
     */
    public function stopListId(Operator $operator, string $value): Session
    {
        $this->setFilter($operator, $value, 'stopListId');
        return $this;
    }

    /**
     * ID страны посетителя
     *
     * @param Operator $operator
     * @param string $value
     * @return Session
     */
    public function countryId(Operator $operator, string $value): Session
    {
        $this->setFilter($operator, $value, 'countryId');
        return $this;
    }

    /**
     * Наименование страны посетителя
     *
     * @param Operator $operator
     * @param string $value
     * @return Session
     */
    public function country(Operator $operator, string $value): Session
    {
        $this->setFilter($operator, $value, 'country');
        return $this;
    }

    /**
     * IP адрес посетителя на последнем хите сессии
     *
     * @param Operator $operator
     * @param string $value
     * @return Session
     */
    public function ip(Operator $operator, string $value): Session
    {
        $this->setFilter($operator, $value, 'ip');
        return $this;
    }

    /**
     * UserAgent посетителя
     *
     * @param Operator $operator
     * @param string $value
     * @return Session
     */
    public function userAgent(Operator $operator, string $value): Session
    {
        $this->setFilter($operator, $value, 'userAgent');
        return $this;
    }

    /**
     * Значение интервала для поля "первого хита сессии"
     *
     * @param Operator $operator
     * @param string $value
     * @return Session
     */
    public function date(Operator $operator, string $value): Session
    {
        $this->setFilter($operator, $value, 'date');
        return $this;
    }

    /**
     * Первая страница сессии
     *
     * @param Operator $operator
     * @param string $value
     * @return Session
     */
    public function urlTo(Operator $operator, string $value): Session
    {
        $this->setFilter($operator, $value, 'urlTo');
        return $this;
    }

    /**
     * Была ли 404 ошибка на первой страницы сессии
     *
     * @param bool $value
     * @return Session
     */
    public function isUrlTo404(bool $value): Session
    {
        $this->setFilter(Operator::Eq, $value, 'isUrlTo404');
        return $this;
    }

    /**
     * ID сайта на первом хите сессии
     *
     * @param Operator $operator
     * @param string $value
     * @return Session
     */
    public function firstSiteId(Operator $operator, string $value): Session
    {
        $this->setFilter($operator, $value, 'firstSiteId');
        return $this;
    }


    /**
     * Последняя страница сессии
     *
     * @param Operator $operator
     * @param string $value
     * @return Session
     */
    public function urlLast(Operator $operator, string $value): Session
    {
        $this->setFilter($operator, $value, 'urlLast');
        return $this;
    }


    /**
     * Была ли 404 ошибка на последней страницы сессии
     *
     * @param bool $value
     * @return Session
     */
    public function isUrlLast404(bool $value): Session
    {
        $this->setFilter(Operator::Eq, $value, 'isUrlLast404');
        return $this;
    }

    /**
     * ID сайта на последнем хите сессии
     *
     * @param Operator $operator
     * @param string $value
     * @return Session
     */
    public function lastSiteId(Operator $operator, string $value): Session
    {
        $this->setFilter($operator, $value, 'lastSiteId');
        return $this;
    }
}