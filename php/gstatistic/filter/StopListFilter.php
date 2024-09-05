<?php

class StopListFilter extends BaseFilter
{
ID* - ID записи стоп-листа;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
DATE_START_1 - начальное значение интервала для поля "время начала активности записи";
DATE_START_2 - конечное значение интервала для поля "время начала активности записи";
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
DATE_END_1 - начальное значение интервала для поля "время окончания активности записи";
DATE_END_2 - конечное значение интервала для поля "время окончания активности записи";
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
ACTIVE - флаг активности записи, воможные значения:
Y - запись активна;
N - запись не активна.
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
SAVE_STATISTIC - флаг необходимости сохранения статистики по посетителю попавшему в стоп-лист, воможные значения:
Y - статистику сохранять;
N - статистику не сохранять.
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
IP_1* - октет 1 IP адреса;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
IP_2* - октет 2 IP адреса;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
IP_3* - октет 3 IP адреса;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
IP_4* - октет 4 IP адреса;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
URL_FROM* - ссылающаяся страница, с которой приходит посетитель;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
USER_AGENT* - UserAgent посетителя;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
MESSAGE* - текст сообщения которое будет выдано посетителю сайта, в случае его попадания под данную запись стоп-листа;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
COMMENTS* - административный комментарий, используется как правило для указания причин создания данной записи;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
URL_TO* - страница (или ее часть) на которую приходит посетитель;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
URL_REDIRECT* - страница на которую необходимо перенаправить посетителя после его попадания под данную запись стоп-листа;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
SITE_ID* - ID сайта для которого запись будет действительна; если значение не задано, то это означает что запись действительная для всех сайтов;
    public function pathId(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }

}