<?php


class StatEventTypeFilter extends BaseFilter
{
ID* - ID типа события;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
EVENT1* - идентификатор event1 типа события;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
EVENT2* - идентификатор event2 типа события;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
NAME* - название типа события;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
DESCRIPTION* - описание типа события;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
DATE_ENTER_1 - начальное значение интервала для поля "дата первого события данного типа";
DATE_ENTER_2 - конечное значение интервала для поля "дата первого события данного типа";
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
DATE_LAST_1 - начальное значение интервала для поля "дата последнего события данного типа";
DATE_LAST_2 - конечное значение интервала для поля "дата последнего события данного типа";
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
DATE1_PERIOD - начальное значение значение для произвольного периода;
DATE2_PERIOD - конечное значение значение для произвольного периода;

    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
    COUNTER1 - начальное значение интервала для поля "суммарное количество событий данного типа";
COUNTER2 - конечное значение интервала для поля "суммарное количество событий данного типа";
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
ADV_VISIBLE - флаг включать ли статистику по данному типу события в отчет по рекламным кампаниям, возможные значения:
Y - включать;
N - не включать.
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
DIAGRAM_DEFAULT - флаг включать ли данный тип события в круговую диаграмму и график по умолчанию, возможные значения:
Y - включать;
N - не включать.
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
KEEP_DAYS1 - начальное значение интервала для поля "количество дней отведенное для хранения событий данного типа";
KEEP_DAYS2 - конечное значение интервала для поля "количество дней отведенное для хранения событий данного типа";
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
DYNAMIC_KEEP_DAYS1 - начальное значение интервала для поля "количество дней отведенное для хранения статистики по данному типу события в разрезе по дням";
DYNAMIC_KEEP_DAYS2 - конечное значение интервала для поля "количество дней отведенное для хранения статистики по данному типу события в разрезе по дням";
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
MONEY1 - начальное значение интервала для поля "суммарная денежная сумма для данного типа событий";
MONEY2 - конечное значение интервала для поля "суммарная денежная сумма для данного типа событий";
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
CURRENCY - трехсимвольный идентификатор валюты для денежной суммы;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
    GROUP - группировка результирующего списка, возможные значения:
event1 - группировка по event1;
event2 - группировка по event2.

}