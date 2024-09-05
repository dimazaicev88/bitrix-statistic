<?php

class PageFilter extends BaseFilter
{

    /**
     * Значение для интервала даты за которую необходимо получить данные
     *
     * @param Operator $operator
     * @param string $date
     * @return PageFilter
     */
    public function date(Operator $operator, string $date): PageFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }

    /**
     * Флаг "показывать только каталоги или только страницы", возможные значения:
     * true - в результирующем списке должны быть только каталоги;
     * false - в результирующем списке должны быть только страницы.
     * @param Operator $operator
     * @param bool $isDir
     * @return void
     */
    public function isDir(Operator $operator, bool $isDir): PageFilter
    {
        $this->setFilter($operator, $isDir, 'isDir');
        return $this;
    }

    /**
     * Полный путь к странице (каталогу) для которой необходимо вывести данные
     *
     * @param Operator $operator
     * @param string $url
     * @return void
     */
    public function url(Operator $operator, string $url): PageFilter
    {
        $this->setFilter($operator, $url, 'url');
        return $this;
    }

    /**
     * Была ли 404 ошибка на странице, возможные значения:
     *
     * @param Operator $operator
     * @param bool $isUrl404
     * @return PageFilter
     */
    public function isUrl404(Operator $operator, bool $isUrl404): PageFilter
    {
        $this->setFilter($operator, $isUrl404, 'isUrl404');
        return $this;
    }

    /**
     * UUID рекламной кампании (РК), данное поле позволяет отфильтровать только те страницы (каталоги) которые были открыты
     * только посетителями по данной РК и соответственно получить данные по посещаемости страницы (каталога) url только этих посетителей;
     *
     * @param Operator $operator
     * @param bool $advUuid
     * @return PageFilter
     */
    public function advUuid(Operator $operator, bool $advUuid): PageFilter
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
     * @return PageFilter
     */
    public function advDataType(Operator $operator, AdvDataType $advDataType): PageFilter
    {
        $this->setFilter($operator, $advDataType->value, 'advDataType');
        return $this;
    }

    /**
     * ID сайта
     *
     * @param Operator $operator
     * @param bool $siteId
     * @return PageFilter
     */
    public function siteId(Operator $operator, bool $siteId): PageFilter
    {
        $this->setFilter($operator, $siteId, 'siteId');
        return $this;
    }
}