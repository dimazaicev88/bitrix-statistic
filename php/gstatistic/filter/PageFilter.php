<?php

class PageFilter extends BaseFilter
{

    /**
     * Значение для интервала даты за которую необходимо получить данные
     *
     * @param Operator $operator
     * @param string $value
     * @return PageFilter
     */
    public function date(Operator $operator, string $value): PageFilter
    {
        $this->setFilter($operator, $value, 'date');
        return $this;
    }

    /**
     * Флаг "показывать только каталоги или только страницы", возможные значения:
     * true - в результирующем списке должны быть только каталоги;
     * false - в результирующем списке должны быть только страницы.
     * @param bool $value
     * @return PageFilter
     */
    public function isDir(bool $value): PageFilter
    {
        $this->setFilter(Operator::Eq, $value, 'isDir');
        return $this;
    }

    /**
     * Полный путь к странице (каталогу) для которой необходимо вывести данные
     *
     * @param Operator $operator
     * @param string $value
     * @return PageFilter
     */
    public function url(Operator $operator, string $value): PageFilter
    {
        $this->setFilter($operator, $value, 'url');
        return $this;
    }

    /**
     * Была ли 404 ошибка на странице, возможные значения:
     *
     * @param bool $value
     * @return PageFilter
     */
    public function isUrl404(bool $value): PageFilter
    {
        $this->setFilter(Operator::Eq, $value, 'isUrl404');
        return $this;
    }

    /**
     * UUID рекламной кампании (РК), данное поле позволяет отфильтровать только те страницы (каталоги) которые были открыты
     * только посетителями по данной РК и соответственно получить данные по посещаемости страницы (каталога) url только этих посетителей;
     *
     * @param Operator $operator
     * @param bool $value
     * @return PageFilter
     */
    public function advUuid(Operator $operator, bool $value): PageFilter
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
     * @return PageFilter
     */
    public function advDataType(Operator $operator, AdvDataType $value): PageFilter
    {
        $this->setFilter($operator, $value->value, 'advDataType');
        return $this;
    }

    /**
     * ID сайта
     *
     * @param Operator $operator
     * @param bool $value
     * @return PageFilter
     */
    public function siteId(Operator $operator, bool $value): PageFilter
    {
        $this->setFilter($operator, $value, 'siteId');
        return $this;
    }
}