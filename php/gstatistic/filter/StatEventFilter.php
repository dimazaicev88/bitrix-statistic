<?php


class StatEventFilter extends BaseFilter
{
ID* - ID события;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
ID_EXACT_MATCH - если значение равно "N", то при фильтрации по ID будет искаться вхождение;
EVENT_ID* - ID типа события;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
EVENT_ID_EXACT_MATCH - если значение равно "N", то при фильтрации по EVENT_ID будет искаться вхождение;
EVENT_NAME* - название типа события;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
EVENT_NAME_EXACT_MATCH - если значение равно "Y", то при фильтрации по EVENT_NAME будет искаться точное совпадение;
EVENT1* - идентификатор event1 типа события;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
EVENT1_EXACT_MATCH - если значение равно "Y", то при фильтрации по EVENT1 будет искаться точное совпадение;
EVENT2* - идентификатор event2 типа события;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
EVENT2_EXACT_MATCH - если значение равно "Y", то при фильтрации по EVENT2 будет искаться точное совпадение;
EVENT3* - дополнительный параметр event3 события;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
EVENT3_EXACT_MATCH - если значение равно "Y", то при фильтрации по EVENT3 будет искаться точное совпадение;
DATE - время события (точное совпадение);
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
DATE1 - начальное значение интервала для поля "дата события";
DATE2 - начальное значение интервала для поля "дата события";
MONEY - денежная сумма события (точное совпадение);
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
MONEY1 - начальное значение интервала для поля "денежная сумма";
MONEY2 - конечное значение интервала для поля "денежная сумма";
CURRENCY - трехсимвольный идентификатор валюты для денежной суммы;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
SESSION_ID* - ID сессии;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
SESSION_ID_EXACT_MATCH - если значение равно "N", то при фильтрации по SESSION_ID будет искаться вхождение;
GUEST_ID* - ID посетителя;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
GUEST_ID_EXACT_MATCH - если значение равно "N", то при фильтрации по GUEST_ID будет искаться вхождение;
ADV_ID* - ID рекламной кампании;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
ADV_BACK - флаг "возврат по рекламной кампании", возможные значения:
Y - был возврат;
N - был прямой заход.
HIT_ID* - ID хита;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
HIT_ID_EXACT_MATCH - если значение равно "N", то при фильтрации по HIT_ID будет искаться вхождение;
COUNTRY_ID* - ID страны посетителя сгенерировавшего событие;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
COUNTRY_ID_EXACT_MATCH - если значение равно "N", то при фильтрации по COUNTRY_ID будет искаться вхождение;
COUNTRY* - название страны посетителя сгенерировавшего событие;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
COUNTRY_EXACT_MATCH - если значение равно "Y", то при фильтрации по COUNTRY будет искаться точное совпадение;
REFERER_URL* - ссылающаяся страница;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
REFERER_URL_EXACT_MATCH - если значение равно "Y", то при фильтрации по REFERER_URL будет искаться точное совпадение;
REFERER_SITE_ID - ID сайта для ссылающейся страницы;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
URL* - страница на которой было зафиксировано событие;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
URL_EXACT_MATCH - если значение равно "Y", то при фильтрации по URL будет искаться точное совпадение;
SITE_ID - ID сайта для страницы на которой было зафиксировано событие;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
REDIRECT_URL* - страница куда был перенаправлен посетитель после фиксации события;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }

}