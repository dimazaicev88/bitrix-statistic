<?php


class StatEventFilter extends BaseFilter
{

    /**
     * UUID события
     *
     * @param Operator $operator
     * @param string $value
     * @return StatEventFilter
     */
    public function uuid(Operator $operator, string $value): StatEventFilter
    {
        $this->setFilter($operator, $value, 'uuid');
        return $this;
    }

    /**
     * UUID типа события
     *
     * @param Operator $operator
     * @param string $value
     * @return StatEventFilter
     */
    public function eventUuid(Operator $operator, string $value): StatEventFilter
    {
        $this->setFilter($operator, $value, 'eventUuid');
        return $this;
    }

    /**
     *Название типа события
     *
     * @param Operator $operator
     * @param string $value
     * @return StatEventFilter
     */
    public function eventName(Operator $operator, string $value): StatEventFilter
    {
        $this->setFilter($operator, $value, 'eventName');
        return $this;
    }


    /**
     * Идентификатор event1 типа события
     *
     * @param Operator $operator
     * @param string $value
     * @return StatEventFilter
     */
    public function event1(Operator $operator, string $value): StatEventFilter
    {
        $this->setFilter($operator, $value, 'event1');
        return $this;
    }


    /**
     * Идентификатор event2 типа события
     *
     * @param Operator $operator
     * @param string $value
     * @return StatEventFilter
     */
    public function event2(Operator $operator, string $value): StatEventFilter
    {
        $this->setFilter($operator, $value, 'event2');
        return $this;
    }

    /**
     * Дополнительный параметр event3 события
     *
     * @param Operator $operator
     * @param string $value
     * @return StatEventFilter
     */
    public function event3(Operator $operator, string $value): StatEventFilter
    {
        $this->setFilter($operator, $value, 'event3');
        return $this;
    }

    /**
     * Время события
     *
     * @param Operator $operator
     * @param string $value
     * @return StatEventFilter
     */
    public function date(Operator $operator, string $value): StatEventFilter
    {
        $this->setFilter($operator, $value, 'date');
        return $this;
    }

    /**
     * Денежная сумма события
     *
     * @param Operator $operator
     * @param string $value
     * @return StatEventFilter
     */
    public function money(Operator $operator, string $value): StatEventFilter
    {
        $this->setFilter($operator, $value, 'money');
        return $this;
    }

    /**
     * Трех символьный идентификатор валюты для денежной суммы
     *
     * @param Operator $operator
     * @param string $value
     * @return StatEventFilter
     */
    public function currency(Operator $operator, string $value): StatEventFilter
    {
        $this->setFilter($operator, $value, 'currency');
        return $this;
    }

    /**
     * UUID сессии
     *
     * @param Operator $operator
     * @param string $value
     * @return StatEventFilter
     */
    public function sessionUuid(Operator $operator, string $value): StatEventFilter
    {
        $this->setFilter($operator, $value, 'sessionUuid');
        return $this;
    }

    /**
     * UUID посетителя
     *
     * @param Operator $operator
     * @param string $value
     * @return StatEventFilter
     */
    public function guestUuid(Operator $operator, string $value): StatEventFilter
    {
        $this->setFilter($operator, $value, 'guestUuid');
        return $this;
    }

    /**
     * UUID рекламной кампании
     *
     * @param Operator $operator
     * @param string $value
     * @return StatEventFilter
     */
    public function advUuid(Operator $operator, string $value): StatEventFilter
    {
        $this->setFilter($operator, $value, 'advUuid');
        return $this;
    }

    /**
     * Флаг "возврат по рекламной кампании"
     *
     * @param Operator $operator
     * @param string $value
     * @return StatEventFilter
     */
    public function advBack(Operator $operator, string $value): StatEventFilter
    {
        $this->setFilter($operator, $value, 'advBack');
        return $this;
    }

    /**
     * UUID хита
     *
     * @param Operator $operator
     * @param string $value
     * @return StatEventFilter
     */
    public function hitUuid(Operator $operator, string $value): StatEventFilter
    {
        $this->setFilter($operator, $value, 'hitUuid');
        return $this;
    }

    /**
     * ID страны посетителя сгенерировавшего событие
     *
     * @param Operator $operator
     * @param string $value
     * @return StatEventFilter
     */
    public function countryId(Operator $operator, string $value): StatEventFilter
    {
        $this->setFilter($operator, $value, 'countryId');
        return $this;
    }

    /**
     * Название страны посетителя сгенерировавшего событие
     *
     * @param Operator $operator
     * @param string $value
     * @return StatEventFilter
     */
    public function country(Operator $operator, string $value): StatEventFilter
    {
        $this->setFilter($operator, $value, 'country');
        return $this;
    }

    /**
     * Ссылающаяся страница
     *
     * @param Operator $operator
     * @param string $value
     * @return StatEventFilter
     */
    public function refererUrl(Operator $operator, string $value): StatEventFilter
    {
        $this->setFilter($operator, $value, 'refererUrl');
        return $this;
    }

    /**
     * ID сайта для ссылающейся страницы
     *
     * @param Operator $operator
     * @param string $value
     * @return StatEventFilter
     */
    public function refererSiteId(Operator $operator, string $value): StatEventFilter
    {
        $this->setFilter($operator, $value, 'refererSiteId');
        return $this;
    }

    /**
     * Страница на которой было зафиксировано событие
     *
     * @param Operator $operator
     * @param string $value
     * @return StatEventFilter
     */
    public function url(Operator $operator, string $value): StatEventFilter
    {
        $this->setFilter($operator, $value, 'url');
        return $this;
    }

    /**
     * ID сайта для страницы на которой было зафиксировано событие
     *
     * @param Operator $operator
     * @param string $value
     * @return StatEventFilter
     */
    public function siteId(Operator $operator, string $value): StatEventFilter
    {
        $this->setFilter($operator, $value, 'siteId');
        return $this;
    }

    /**
     *  Страница куда был перенаправлен посетитель после фиксации события
     * @param Operator $operator
     * @param string $value
     * @return StatEventFilter
     */
    public function redirectUrl(Operator $operator, string $value): StatEventFilter
    {
        $this->setFilter($operator, $value, 'redirectUrl');
        return $this;
    }

}