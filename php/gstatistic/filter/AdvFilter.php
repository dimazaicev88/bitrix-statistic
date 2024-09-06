<?php

class AdvFilter extends BaseFilter
{

    /**
     * Referer1
     *
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    function referer1(Operator $operator, string $value): AdvFilter
    {
        $this->setFilter($operator, $value, "referer1");
        return $this;
    }

    /**
     * Referer2
     *
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    function referer2(Operator $operator, string $value): AdvFilter
    {
        $this->setFilter($operator, $value, "referer2");
        return $this;
    }


    /**
     * Посетителей на прямом заходе
     *
     * @param Operator $operator
     * @param int $value
     * @return $this
     */
    function guests(Operator $operator, int $value): AdvFilter
    {
        $this->setFilter($operator, $value, "guests");
        return $this;
    }

    /**
     * Посетителей на возврате
     *
     * @param Operator $operator
     * @param int $value
     * @return $this
     */
    function guestsBack(Operator $operator, int $value): AdvFilter
    {
        $this->setFilter($operator, $value, "guestsBack");
        return $this;
    }

    /**
     * Новых посетителей
     *
     * @param Operator $operator
     * @param int $value
     * @return $this
     */
    function newGuests(Operator $operator, int $value): AdvFilter
    {
        $this->setFilter($operator, $value, "newGuests");
        return $this;
    }

    /**
     * Посетителей, добавивших сайт в "Избранное" на прямом заходе
     *
     * @param Operator $operator
     * @param int $value
     * @return $this
     */
    function favorites(Operator $operator, int $value): AdvFilter
    {
        $this->setFilter($operator, $value, "favorites");
        return $this;
    }

    /**
     * Посетителей, добавившие сайт в "Избранное" на возврате
     *
     * @param Operator $operator
     * @param int $value
     * @return $this
     */
    function favoritesBack(Operator $operator, int $value): AdvFilter
    {
        $this->setFilter($operator, $value, "favoritesBack");
        return $this;
    }

    /**
     * Хостов на прямом заходе
     *
     * @param Operator $operator
     * @param int $value
     * @return $this
     */
    function hosts(Operator $operator, int $value): AdvFilter
    {
        $this->setFilter($operator, $value, "hosts");
        return $this;
    }

    /**
     * Хостов на возврате
     *
     * @param Operator $operator
     * @param int $value
     * @return $this
     */
    function hostsBack(Operator $operator, int $value): AdvFilter
    {
        $this->setFilter($operator, $value, "hostsBack");
        return $this;
    }

    /**
     * Сессий на прямом заходе
     *
     * @param Operator $operator
     * @param int $value
     * @return $this
     */
    function sessions(Operator $operator, int $value): AdvFilter
    {
        $this->setFilter($operator, $value, "sessions");
        return $this;
    }

    /**
     * Сессий на возврате
     *
     * @param Operator $operator
     * @param int $value
     * @return $this
     */
    function sessionsBack(Operator $operator, int $value): AdvFilter
    {
        $this->setFilter($operator, $value, "sessionsBack");
        return $this;
    }

    /**
     * Хитов на прямом
     *
     * @param Operator $operator
     * @param int $value
     * @return $this
     */
    function hits(Operator $operator, int $value): AdvFilter
    {
        $this->setFilter($operator, $value, "hits");
        return $this;
    }

    /**
     * Хитов на возврате
     *
     * @param Operator $operator
     * @param int $value
     * @return $this
     */
    function hitsBack(Operator $operator, int $value): AdvFilter
    {
        $this->setFilter($operator, $value, "hitsBack");
        return $this;
    }

    /**
     * Дата для фильтрации
     *
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    function period(Operator $operator, string $value): AdvFilter
    {
        $this->setFilter($operator, $value, "period");
        return $this;
    }
}