<?php

class RefererFilter extends BaseFilter
{

Массив для фильтрации результирующего списка. В массиве допустимы следующие ключи:
ID - ID записи;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
SESSION_ID - ID сессии;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }

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
FROM_PROTOCOL - протокол ссылающейся страницы;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
FROM_DOMAIN - домен ссылающейся страницы;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
FROM_PAGE - ссылающаяся страница;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
FROM - протокол + домен + ссылающаяся страница;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
TO* - страница на которую пришли;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
TO_404 - была ли 404 ошибка на странице, на которую пришли:
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }

SITE_ID - ID сайта на который пришли;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
GROUP - группировка результирующего списка; возможные значения:
S - группировка по ссылающемуся домену (сайту);
U - группировка по ссылающейся странице.

}