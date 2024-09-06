<?php

class RefererFilter extends BaseFilter
{

    /**
     * ID записи;
     *
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    public function uuid(Operator $operator, string $value): RefererFilter
    {
        $this->setFilter($operator, $value, 'uuid');
        return $this;
    }

    /**
     * UUID сессии;
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    public function sessionUuid(Operator $operator, string $value): RefererFilter
    {
        $this->setFilter($operator, $value, 'sessionUuid');
        return $this;
    }

    /**
     * Значение интервала для поля "дата"
     *
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    public function date(Operator $operator, string $value): RefererFilter
    {
        $this->setFilter($operator, $value, 'date');
        return $this;
    }

    /**
     * Протокол ссылающейся страницы;
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    public function fromProtocol(Operator $operator, string $value): RefererFilter
    {
        $this->setFilter($operator, $value, 'fromProtocol');
        return $this;
    }

    /**
     * Домен ссылающейся страницы;
     *
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    public function fromDomain(Operator $operator, string $value): RefererFilter
    {
        $this->setFilter($operator, $value, 'fromDomain');
        return $this;
    }

    /**
     * Ссылающаяся страница;
     *
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    public function fromPage(Operator $operator, string $value): RefererFilter
    {
        $this->setFilter($operator, $value, 'fromPage');
        return $this;
    }

    /**
     * Протокол + домен + ссылающаяся страница;
     *
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    public function from(Operator $operator, string $value): RefererFilter
    {
        $this->setFilter($operator, $value, 'from');
        return $this;
    }

    /**
     *Страница на которую пришли
     *
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    public function to(Operator $operator, string $value): RefererFilter
    {
        $this->setFilter($operator, $value, 'to');
        return $this;
    }

    /**
     * Была ли 404 ошибка на странице, на которую пришли
     *
     * @param bool $value
     * @return $this
     */
    public function to404(bool $value): RefererFilter
    {
        $this->setFilter(Operator::Eq, $value, 'to404');
        return $this;
    }


    /**
     * ID сайта на который пришли
     *
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    public function siteId(Operator $operator, string $value): RefererFilter
    {
        $this->setFilter($operator, $value, 'siteId');
        return $this;
    }

//GROUP - группировка результирующего списка;
//возможные значения:
//S - группировка по ссылающемуся домену (сайту);
//U - группировка по ссылающейся странице.

}