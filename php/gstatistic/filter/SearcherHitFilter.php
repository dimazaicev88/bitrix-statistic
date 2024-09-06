<?php

class SearcherHitFilter extends BaseFilter
{

    /**
     * UUID хита;
     *
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    public function hitUuid(Operator $operator, string $value): SearcherHitFilter
    {
        $this->setFilter($operator, $value, 'hitUuid');
        return $this;
    }

    /**
     * UUID поисковой системы;
     *
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    public function searcherUuid(Operator $operator, string $value): SearcherHitFilter
    {
        $this->setFilter($operator, $value, 'searcherUuid');
        return $this;
    }


    /**
     * URL* - адрес проиндексированной страницы;
     *
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    public function url(Operator $operator, string $value): SearcherHitFilter
    {
        $this->setFilter($operator, $value, 'date');
        return $this;
    }

    /**
     * Была ли 404 ошибка на проиндексированной странице
     *
     * @param bool $date
     * @return $this
     */
    public function isUrl404(bool $date): SearcherHitFilter
    {
        $this->setFilter(Operator::Eq, $date, 'date');
        return $this;
    }

    /**
     * Название поисковой системы;
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    public function searcher(Operator $operator, string $value): SearcherHitFilter
    {
        $this->setFilter($operator, $value, 'searcher');
        return $this;
    }

    /**
     * Значение интервала для поля "дата"
     *
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    public function date(Operator $operator, string $value): SearcherHitFilter
    {
        $this->setFilter($operator, $value, 'date');
        return $this;
    }

    /**
     * IP адрес поисковой системы
     *
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    public function ip(Operator $operator, string $value): SearcherHitFilter
    {
        $this->setFilter($operator, $value, 'date');
        return $this;
    }

    /**
     * UserAgent поисковой системы
     *
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    public function userAgent(Operator $operator, string $value): SearcherHitFilter
    {
        $this->setFilter($operator, $value, 'userAgent');
        return $this;
    }

    /**
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    public function siteId(Operator $operator, string $value): SearcherHitFilter
    {
        $this->setFilter($operator, $value, 'siteId');
        return $this;
    }
}