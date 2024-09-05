<?php

class HitFilter extends BaseFilter
{

    /**
     * ID - UUID хита;
     * @param Operator $operator
     * @param string $uuid
     * @return HitFilter
     */
    function uuid(Operator $operator, string $uuid): HitFilter
    {
        $this->setFilter($operator, $uuid, 'uuid');
        return $this;
    }

    /**
     * Guest UUID посетителя
     *
     * @param Operator $operator
     * @param string $uuid
     * @return HitFilter
     */
    function guestUuid(Operator $operator, string $uuid): HitFilter
    {
        $this->setFilter($operator, $uuid, 'guestUuid');
        return $this;
    }

    /**
     * Флаг "новый посетитель"
     *
     * @param bool $value
     * @return HitFilter
     */
    function isNewGuest(bool $value): HitFilter
    {
        $this->setFilter(Operator::Eq, $value, 'isNewGuest');
        return $this;
    }

    /**
     * UUID сессии
     * @param Operator $operator
     * @param string $sessionUuid
     * @return HitFilter
     */
    function sessionUuid(Operator $operator, string $sessionUuid): HitFilter
    {
        $this->setFilter($operator, $sessionUuid, 'sessionUuid');
        return $this;
    }

    /**
     * UUID записи стоп - листа под которую попал посетитель(если это имело место)
     *
     * @param Operator $operator
     * @param string $stopListUuid
     * @return HitFilter
     */
    function stopListUuid(Operator $operator, string $stopListUuid): HitFilter
    {
        $this->setFilter($operator, $stopListUuid, 'stopListUuid');
        return $this;
    }

    /**
     * Страница хита
     *
     * @param Operator $operator
     * @param string $url
     * @return HitFilter
     */
    function url(Operator $operator, string $url): HitFilter
    {
        $this->setFilter($operator, $url, 'url');
        return $this;
    }

    /**
     * Была ли 404 ошибка на странице хита
     *
     * @param bool $isUrl404
     * @return HitFilter
     */
    function isUrl404(bool $isUrl404): HitFilter
    {
        $this->setFilter(Operator::Eq, $isUrl404, 'isUrl404');
        return $this;
    }

    /**
     * ID пользователя под которым был авторизован посетитель в момент хита или до него;
     *
     * @param Operator $operator
     * @param int $userId
     * @return HitFilter
     */
    function userId(Operator $operator, int $userId): HitFilter
    {
        $this->setFilter($operator, $userId, 'userId');
        return $this;
    }

    /**
     * Флаг "был ли авторизован посетитель в момент хита или до этого"
     *
     * @param bool $isRegistered
     * @return HitFilter
     */
    function isRegistered(bool $isRegistered): HitFilter
    {
        $this->setFilter(Operator::Eq, $isRegistered, 'isRegistered');
        return $this;
    }

    /**
     * Значение интервала даты хита
     *
     * @param Operator $operator
     * @param string $date
     * @return HitFilter
     */
    function date(Operator $operator, string $date): HitFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }

    /**
     * IP адрес посетителя в момент хита
     *
     * @param Operator $operator
     * @param string $ip
     * @return HitFilter
     */
    function ip(Operator $operator, string $ip): HitFilter
    {
        $this->setFilter($operator, $ip, 'ip');
        return $this;
    }

    /**
     * UserAgent посетителя в момент хита
     *
     * @param Operator $operator
     * @param string $userAgent
     * @return HitFilter
     */
    function userAgent(Operator $operator, string $userAgent): HitFilter
    {
        $this->setFilter($operator, $userAgent, 'userAgent');
        return $this;
    }

    /**
     * ID страны посетителя в момент хита
     *
     * @param Operator $operator
     * @param string $countryId
     * @return HitFilter
     */
    function countryId(Operator $operator, string $countryId): HitFilter
    {
        $this->setFilter($operator, $countryId, 'countryId');
        return $this;
    }

    /**
     * Название страны
     *
     * @param Operator $operator
     * @param string $country
     * @return HitFilter
     */
    function country(Operator $operator, string $country): HitFilter
    {
        $this->setFilter($operator, $country, 'country');
        return $this;
    }

    /**
     * Содержимое Cookie в момент хита
     *
     * @param Operator $operator
     * @param string $cookie
     * @return HitFilter
     */
    function cookie(Operator $operator, string $cookie): HitFilter
    {
        $this->setFilter($operator, $cookie, 'cookie');
        return $this;
    }

    /**
     * Содержимое Cookie в момент хита
     *
     * @param bool $isStop
     * @return HitFilter
     */
    function isStop(bool $isStop): HitFilter
    {
        $this->setFilter(Operator::Eq, $isStop, 'isStop');
        return $this;
    }

    /**
     * ID сайта
     *
     * @param Operator $operator
     * @param bool $siteId
     * @return HitFilter
     */
    function siteId(Operator $operator, bool $siteId): HitFilter
    {
        $this->setFilter($operator, $siteId, 'siteId');
        return $this;
    }
}