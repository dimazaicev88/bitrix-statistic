<?php

class PathFilter extends BaseFilter
{

    /**
     * ID отрезка пути;
     *
     * @param Operator $operator
     * @param string $value
     * @return PathFilter
     */
    public function pathId(Operator $operator, string $value): PathFilter
    {
        $this->setFilter($operator, $value, 'pathId');
        return $this;
    }

    /**
     * Значение для интервала даты
     *
     * @param Operator $operator
     * @param string $value
     * @return PathFilter
     */
    public function date(Operator $operator, string $value): PathFilter
    {
        $this->setFilter($operator, $value, 'date');
        return $this;
    }

    /**
     * Первая страница пути;
     *
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    public function firstPage(Operator $operator, string $value): PathFilter
    {
        $this->setFilter($operator, $value, 'firstPage');
        return $this;
    }

    /**
     * ID сайта первой страницы пути;
     *
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    public function firstPageSiteId(Operator $operator, string $value): PathFilter
    {
        $this->setFilter($operator, $value, 'firstPageSiteId');
        return $this;
    }

    /**
     * Была ли 404 ошибка на первой странице пути
     *
     * @param Operator $operator
     * @param bool $value
     * @return $this
     */
    public function isFirstPage404(Operator $operator, bool $value): PathFilter
    {
        $this->setFilter($operator, $value, 'isFirstPage404');
        return $this;
    }


    /**
     * Последняя страница пути;
     *
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    public function lastPage(Operator $operator, string $value): PathFilter
    {
        $this->setFilter($operator, $value, 'lastPage');
        return $this;
    }

    /**
     * ID сайта последней страницы пути;
     *
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    public function lastPageSiteId(Operator $operator, string $value): PathFilter
    {
        $this->setFilter($operator, $value, 'lastPageSiteId');
        return $this;
    }

    /**
     * Была ли 404 ошибка на последней странице пути
     *
     * @param bool $isLastPage404
     * @return $this
     */
    public function isLastPage404(bool $isLastPage404): PathFilter
    {
        $this->setFilter(Operator::Eq, $isLastPage404, 'isLastPage404');
        return $this;
    }

    /**
     * Произвольная страница пути
     *
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    public function page(Operator $operator, string $value): PathFilter
    {
        $this->setFilter($operator, $value, 'page');
        return $this;
    }

    /**
     * ID сайта произвольной страницы пути
     *
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    public function pageSiteId(Operator $operator, string $value): PathFilter
    {
        $this->setFilter($operator, $value, 'pageSiteId');
        return $this;
    }

    /**
     * Была ли 404 ошибка на произвольной странице пути
     *
     * @param bool $value
     * @return $this
     */
    public function isPage404(bool $value): PathFilter
    {
        $this->setFilter(Operator::Eq, $value, 'isPage404');
        return $this;
    }

    /**
     * UUID рекламной кампании, по посетителям которой надо получить данные;
     *
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    public function advUuid(Operator $operator, string $value): PathFilter
    {
        $this->setFilter($operator, $value, 'advUuid');
        return $this;
    }

    /**
     * Флаг типа данных для рекламной кампании, возможные значения:
     *
     * P - только по прямым заходам по рекламной кампании;
     * B - только по возвратам по рекламной кампании;
     * S - сумма по прямым заходам и возвратам.
     *
     * @param Operator $operator
     * @param AdvDataType $value
     * @return $this
     */
    public function advDataType(Operator $operator, AdvDataType $value): PathFilter
    {
        $this->setFilter($operator, $value->value, 'advDataType');
        return $this;
    }

    /**
     * Значение интервала для поля "количество страниц в пути";
     *
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    public function steps(Operator $operator, string $value): PathFilter
    {
        $this->setFilter($operator, $value, 'steps');
        return $this;
    }
}