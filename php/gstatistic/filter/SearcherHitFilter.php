<?php

class SearcherHitFilter extends BaseFilter
{


    /**
     * UUID хита;
     *
     * @param Operator $operator
     * @param string $date
     * @return PathFilter
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
     * @param string $date
     * @return PathFilter
     */
    public function pathId(Operator $operator, string $date): SearcherHitFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }

    URL* - адрес проиндексированной страницы;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }

    URL_404 - была ли 404 ошибка на проиндексированной странице:
    Y - была;
    N - не была.
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }

    SEARCHER* - название поисковой системы;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }

    SEARCHER_EXACT_MATCH - если значение равно "Y", то при фильтрации по SEARCHER будет искаться точное совпадение;
    /**
     * Значение интервала для поля "дата"
     *
     * @param Operator $operator
     * @param string $date
     * @return $this
     */
    public function date(Operator $operator, string $date): PhraseFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }

    IP* - IP адрес поисковой системы;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }

    IP_EXACT_MATCH - если значение равно "Y", то при фильтрации по IP будет искаться точное совпадение;
    USER_AGENT* - UserAgent поисковой системы;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }

    USER_AGENT_EXACT_MATCH - если значение равно "Y", то при фильтрации по USER_AGENT будет искаться точное совпадение;
    SITE_ID* - ID сайта;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
    SITE_ID_EXACT_MATCH - если значение равно "N", то при фильтрации по SITE_ID будет искаться вхождение.

}