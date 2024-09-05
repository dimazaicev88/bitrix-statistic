<?php

class SearcherHitFilter extends BaseFilter
{

    /**
     * UUID хита;
     *
     * @param Operator $operator
     * @param string $hitUuid
     * @return $this
     */
    public function hitUuid(Operator $operator, string $hitUuid): SearcherHitFilter
    {
        $this->setFilter($operator, $hitUuid, 'hitUuid');
        return $this;
    }

    /**
     * UUID поисковой системы;
     *
     * @param Operator $operator
     * @param string $searcherUuid
     * @return $this
     */
    public function searcherUuid(Operator $operator, string $searcherUuid): SearcherHitFilter
    {
        $this->setFilter($operator, $searcherUuid, 'searcherUuid');
        return $this;
    }


    /**
     * URL* - адрес проиндексированной страницы;
     *
     * @param Operator $operator
     * @param string $date
     * @return $this
     */
    public function url(Operator $operator, string $date): SearcherHitFilter
    {
        $this->setFilter($operator, $date, 'date');
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
     * @param string $searcher
     * @return $this
     */
    public function searcher(Operator $operator, string $searcher): SearcherHitFilter
    {
        $this->setFilter($operator, $searcher, 'searcher');
        return $this;
    }

    /**
     * Значение интервала для поля "дата"
     *
     * @param Operator $operator
     * @param string $date
     * @return $this
     */
    public function date(Operator $operator, string $date): SearcherHitFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }

    /**
     * IP адрес поисковой системы
     *
     * @param Operator $operator
     * @param string $date
     * @return $this
     */
    public function ip(Operator $operator, string $date): SearcherHitFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }

    /**
     * UserAgent поисковой системы
     *
     * @param Operator $operator
     * @param string $userAgent
     * @return $this
     */
    public function userAgent(Operator $operator, string $userAgent): SearcherHitFilter
    {
        $this->setFilter($operator, $userAgent, 'userAgent');
        return $this;
    }

    /**
     * @param Operator $operator
     * @param string $siteId
     * @return $this
     */
    public function siteId(Operator $operator, string $siteId): SearcherHitFilter
    {
        $this->setFilter($operator, $siteId, 'siteId');
        return $this;
    }
}