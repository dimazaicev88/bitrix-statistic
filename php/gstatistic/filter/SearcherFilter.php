<?php


class SearcherFilter extends BaseFilter
{

    /**
     * UUID поисковой системы;
     *
     * @param Operator $operator
     * @param string $date
     * @return SearcherFilter
     */
    public function uuid(Operator $operator, string $date): SearcherFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }

    /**
     * Флаг активности:
     *
     * @param Operator $operator
     * @param bool $active
     * @return $this
     */
    public function active(Operator $operator, bool $active): SearcherFilter
    {
        $this->setFilter($operator, $active, 'active');
        return $this;
    }

    /**
     * Флаг "сохранять хиты поисковой системы", возможные значения
     *
     * @param Operator $operator
     * @param bool $saveStatistic
     * @return SearcherFilter
     */
    public function saveStatistic(Operator $operator, bool $saveStatistic): SearcherFilter
    {
        $this->setFilter($operator, $saveStatistic, 'saveStatistic');
        return $this;
    }

    /**
     * Количество хитов;
     *
     * @param Operator $operator
     * @param string $date
     * @return $this
     */
    public function hits(Operator $operator, string $date): SearcherFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }


    /**
     * Начальное значение для произвольного периода
     *
     * @param Operator $operator
     * @param string $period
     * @return $this
     */
    public function period(Operator $operator, string $period): SearcherFilter
    {
        $this->setFilter($operator, $period, 'period');
        return $this;
    }

    /**
     * Наименование поисковой системы;
     *
     * @param Operator $operator
     * @param string $date
     * @return $this
     */
    public function pathId(Operator $operator, string $date): SearcherFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }

    /**
     * UserAgent поисковой системы;
     *
     * @param Operator $operator
     * @param string $userAgent
     * @return $this
     */
    public function userAgent(Operator $operator, string $userAgent): SearcherFilter
    {
        $this->setFilter($operator, $userAgent, 'userAgent');
        return $this;
    }

}