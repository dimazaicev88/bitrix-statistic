<?php

class RefererFilter extends BaseFilter
{

    /**
     * ID записи;
     *
     * @param Operator $operator
     * @param string $date
     * @return $this
     */
    public function pathId(Operator $operator, string $date): RefererFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }

    /**
     * UUID сессии;
     * @param Operator $operator
     * @param string $sessionUuid
     * @return $this
     */
    public function sessionUuid(Operator $operator, string $sessionUuid): RefererFilter
    {
        $this->setFilter($operator, $sessionUuid, 'sessionUuid');
        return $this;
    }

    /**
     * Значение интервала для поля "дата"
     *
     * @param Operator $operator
     * @param string $date
     * @return $this
     */
    public function date(Operator $operator, string $date): RefererFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }

    /**
     * Протокол ссылающейся страницы;
     * @param Operator $operator
     * @param string $fromProtocol
     * @return $this
     */
    public function fromProtocol(Operator $operator, string $fromProtocol): RefererFilter
    {
        $this->setFilter($operator, $fromProtocol, 'fromProtocol');
        return $this;
    }

    /**
     * Домен ссылающейся страницы;
     *
     * @param Operator $operator
     * @param string $fromDomain
     * @return $this
     */
    public function fromDomain(Operator $operator, string $fromDomain): RefererFilter
    {
        $this->setFilter($operator, $fromDomain, 'fromDomain');
        return $this;
    }

    /**
     * Ссылающаяся страница;
     *
     * @param Operator $operator
     * @param string $fromPage
     * @return $this
     */
    public function fromPage(Operator $operator, string $fromPage): RefererFilter
    {
        $this->setFilter($operator, $fromPage, 'fromPage');
        return $this;
    }

    /**
     * Протокол + домен + ссылающаяся страница;
     *
     * @param Operator $operator
     * @param string $from
     * @return $this
     */
    public function from(Operator $operator, string $from): RefererFilter
    {
        $this->setFilter($operator, $from, 'from');
        return $this;
    }

    /**
     *Страница на которую пришли
     *
     * @param Operator $operator
     * @param string $to
     * @return $this
     */
    public function to(Operator $operator, string $to): RefererFilter
    {
        $this->setFilter($operator, $to, 'to');
        return $this;
    }

    /**
     * Была ли 404 ошибка на странице, на которую пришли
     *
     * @param bool $to404
     * @return $this
     */
    public function to404(bool $to404): RefererFilter
    {
        $this->setFilter(Operator::Eq, $to404, 'to404');
        return $this;
    }


    /**
     * ID сайта на который пришли
     *
     * @param Operator $operator
     * @param string $siteId
     * @return $this
     */
    public function siteId(Operator $operator, string $siteId): RefererFilter
    {
        $this->setFilter($operator, $siteId, 'siteId');
        return $this;
    }

//GROUP - группировка результирующего списка;
//возможные значения:
//S - группировка по ссылающемуся домену (сайту);
//U - группировка по ссылающейся странице.

}