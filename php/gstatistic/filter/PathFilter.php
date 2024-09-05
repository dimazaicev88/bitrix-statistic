<?php

class PathFilter extends BaseFilter
{

    /**
     * ID отрезка пути;
     *
     * @param Operator $operator
     * @param string $date
     * @return PathFilter
     */
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }

    /**
     * Значение для интервала даты
     *
     * @param Operator $operator
     * @param string $date
     * @return PathFilter
     */
    public function date(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }

    /**
     * Первая страница пути;
     *
     * @param Operator $operator
     * @param string $firstPage
     * @return $this
     */
    public function firstPage(Operator $operator, string $firstPage): PathFilter
    {
        $this->setFilter($operator, $firstPage, 'firstPage');
        return $this;
    }

    /**
     * ID сайта первой страницы пути;
     *
     * @param Operator $operator
     * @param string $firstPageSiteId
     * @return $this
     */
    public function firstPageSiteId(Operator $operator, string $firstPageSiteId): PathFilter
    {
        $this->setFilter($operator, $firstPageSiteId, 'firstPageSiteId');
        return $this;
    }

    /**
     * Была ли 404 ошибка на первой странице пути
     *
     * @param Operator $operator
     * @param bool $isFirstPage404
     * @return $this
     */
    public function isFirstPage404(Operator $operator, bool $isFirstPage404): PathFilter
    {
        $this->setFilter($operator, $isFirstPage404, 'isFirstPage404');
        return $this;
    }


    /**
     * Последняя страница пути;
     *
     * @param Operator $operator
     * @param string $lastPage
     * @return $this
     */
    public function lastPage(Operator $operator, string $lastPage): PathFilter
    {
        $this->setFilter($operator, $lastPage, 'lastPage');
        return $this;
    }

    /**
     * ID сайта последней страницы пути;
     *
     * @param Operator $operator
     * @param string $lastPageSiteId
     * @return $this
     */
    public function lastPageSiteId(Operator $operator, string $lastPageSiteId): PathFilter
    {
        $this->setFilter($operator, $lastPageSiteId, 'lastPageSiteId');
        return $this;
    }

    /**
     * Была ли 404 ошибка на последней странице пути
     *
     * @param Operator $operator
     * @param string $isLastPage404
     * @return $this
     */
    public function isLastPage404(Operator $operator, string $isLastPage404): PathFilter
    {
        $this->setFilter($operator, $isLastPage404, 'isLastPage404');
        return $this;
    }

    /**
     * Произвольная страница пути
     *
     * @param Operator $operator
     * @param string $page
     * @return $this
     */
    public function page(Operator $operator, string $page): PathFilter
    {
        $this->setFilter($operator, $page, 'page');
        return $this;
    }

    /**
     * ID сайта произвольной страницы пути
     *
     * @param Operator $operator
     * @param string $pageSiteId
     * @return $this
     */
    public function pageSiteId(Operator $operator, string $pageSiteId): PathFilter
    {
        $this->setFilter($operator, $pageSiteId, 'pageSiteId');
        return $this;
    }

    /**
     * Была ли 404 ошибка на произвольной странице пути
     *
     * @param Operator $operator
     * @param string $isPage404
     * @return $this
     */
    public function isPage404(Operator $operator, string $isPage404): PathFilter
    {
        $this->setFilter($operator, $isPage404, 'isPage404');
        return $this;
    }

    /**
     * UUID рекламной кампании, по посетителям которой надо получить данные;
     *
     * @param Operator $operator
     * @param string $advUuid
     * @return $this
     */
    public function advUuid(Operator $operator, string $advUuid): PathFilter
    {
        $this->setFilter($operator, $advUuid, 'advUuid');
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
     * @param AdvDataType $advDataType
     * @return $this
     */
    public function advDataType(Operator $operator, AdvDataType $advDataType): PathFilter
    {
        $this->setFilter($operator, $advDataType->value, 'advDataType');
        return $this;
    }

    /**
     * Значение интервала для поля "количество страниц в пути";
     *
     * @param Operator $operator
     * @param string $steps
     * @return $this
     */
    public function steps(Operator $operator, string $steps): PathFilter
    {
        $this->setFilter($operator, $steps, 'steps');
        return $this;
    }
}