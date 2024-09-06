<?php

class HitFilter extends BaseFilter
{

    /**
     * ID - UUID хита;
     * @param Operator $operator
     * @param string $value
     * @return HitFilter
     */
    function uuid(Operator $operator, string $value): HitFilter
    {
        $this->setFilter($operator, $value, 'uuid');
        return $this;
    }

    /**
     * Guest UUID посетителя
     *
     * @param Operator $operator
     * @param string $value
     * @return HitFilter
     */
    function guestUuid(Operator $operator, string $value): HitFilter
    {
        $this->setFilter($operator, $value, 'guestUuid');
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
     * @param string $value
     * @return HitFilter
     */
    function sessionUuid(Operator $operator, string $value): HitFilter
    {
        $this->setFilter($operator, $value, 'sessionUuid');
        return $this;
    }

    /**
     * UUID записи стоп - листа под которую попал посетитель(если это имело место)
     *
     * @param Operator $operator
     * @param string $value
     * @return HitFilter
     */
    function stopListUuid(Operator $operator, string $value): HitFilter
    {
        $this->setFilter($operator, $value, 'stopListUuid');
        return $this;
    }

    /**
     * Страница хита
     *
     * @param Operator $operator
     * @param string $value
     * @return HitFilter
     */
    function url(Operator $operator, string $value): HitFilter
    {
        $this->setFilter($operator, $value, 'url');
        return $this;
    }

    /**
     * Была ли 404 ошибка на странице хита
     *
     * @param bool $value
     * @return HitFilter
     */
    function isUrl404(bool $value): HitFilter
    {
        $this->setFilter(Operator::Eq, $value, 'isUrl404');
        return $this;
    }

    /**
     * ID пользователя под которым был авторизован посетитель в момент хита или до него;
     *
     * @param Operator $operator
     * @param int $value
     * @return HitFilter
     */
    function userId(Operator $operator, int $value): HitFilter
    {
        $this->setFilter($operator, $value, 'userId');
        return $this;
    }

    /**
     * Флаг "был ли авторизован посетитель в момент хита или до этого"
     *
     * @param bool $value
     * @return HitFilter
     */
    function isRegistered(bool $value): HitFilter
    {
        $this->setFilter(Operator::Eq, $value, 'isRegistered');
        return $this;
    }

    /**
     * Значение интервала даты хита
     *
     * @param Operator $operator
     * @param string $value
     * @return HitFilter
     */
    function date(Operator $operator, string $value): HitFilter
    {
        $this->setFilter($operator, $value, 'date');
        return $this;
    }

    /**
     * IP адрес посетителя в момент хита
     *
     * @param Operator $operator
     * @param string $value
     * @return HitFilter
     */
    function ip(Operator $operator, string $value): HitFilter
    {
        $this->setFilter($operator, $value, 'ip');
        return $this;
    }

    /**
     * UserAgent посетителя в момент хита
     *
     * @param Operator $operator
     * @param string $value
     * @return HitFilter
     */
    function userAgent(Operator $operator, string $value): HitFilter
    {
        $this->setFilter($operator, $value, 'userAgent');
        return $this;
    }

    /**
     * ID страны посетителя в момент хита
     *
     * @param Operator $operator
     * @param string $value
     * @return HitFilter
     */
    function countryId(Operator $operator, string $value): HitFilter
    {
        $this->setFilter($operator, $value, 'countryId');
        return $this;
    }

    /**
     * Название страны
     *
     * @param Operator $operator
     * @param string $value
     * @return HitFilter
     */
    function country(Operator $operator, string $value): HitFilter
    {
        $this->setFilter($operator, $value, 'country');
        return $this;
    }

    /**
     * Содержимое Cookie в момент хита
     *
     * @param Operator $operator
     * @param string $value
     * @return HitFilter
     */
    function cookie(Operator $operator, string $value): HitFilter
    {
        $this->setFilter($operator, $value, 'cookie');
        return $this;
    }

    /**
     * Содержимое Cookie в момент хита
     *
     * @param bool $value
     * @return HitFilter
     */
    function isStop(bool $value): HitFilter
    {
        $this->setFilter(Operator::Eq, $value, 'isStop');
        return $this;
    }

    /**
     * ID сайта
     *
     * @param Operator $operator
     * @param bool $value
     * @return HitFilter
     */
    function siteId(Operator $operator, bool $value): HitFilter
    {
        $this->setFilter($operator, $value, 'siteId');
        return $this;
    }
}