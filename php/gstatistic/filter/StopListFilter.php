<?php

class StopListFilter extends BaseFilter
{

    /**
     * UUID записи стоп-листа
     *
     * @param Operator $operator
     * @param string $date
     * @return PathFilter
     */
    public function uuid(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'uuid');
        return $this;
    }
DATE_START_1 - начальное значение интервала для поля "время начала активности записи";
DATE_START_2 - конечное значение интервала для поля "время начала активности записи";
    /**
     * @param Operator $operator
     * @param string $date
     * @return PathFilter
     */
    public function date_start(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
DATE_END_1 - начальное значение интервала для поля "время окончания активности записи";
DATE_END_2 - конечное значение интервала для поля "время окончания активности записи";
    /**
     * @param Operator $operator
     * @param string $date
     * @return PathFilter
     */
    public function date_end(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
ACTIVE - флаг активности записи, воможные значения:
Y - запись активна;
N - запись не активна.
    /**
     * @param Operator $operator
     * @param string $date
     * @return PathFilter
     */
    public function active(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
SAVE_STATISTIC - флаг необходимости сохранения статистики по посетителю попавшему в стоп-лист, воможные значения:
Y - статистику сохранять;
N - статистику не сохранять.
    /**
     * @param Operator $operator
     * @param string $date
     * @return PathFilter
     */
    public function save_statistic(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
IP_1* - октет 1 IP адреса;

    /**
     * @param Operator $operator
     * @param string $date
     * @return PathFilter
     */
    public function ip_1(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
IP_2* - октет 2 IP адреса;
    /**
     * @param Operator $operator
     * @param string $date
     * @return PathFilter
     */
    public function ip_2(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
IP_3* - октет 3 IP адреса;
    /**
     * @param Operator $operator
     * @param string $date
     * @return PathFilter
     */
    public function ip_3(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
IP_4* - октет 4 IP адреса;
    /**
     * @param Operator $operator
     * @param string $date
     * @return PathFilter
     */
    public function ip_4(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
URL_FROM* - ссылающаяся страница, с которой приходит посетитель;
    /**
     * @param Operator $operator
     * @param string $date
     * @return PathFilter
     */
    public function url_from(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }

USER_AGENT* - UserAgent посетителя;

    /**
     * @param Operator $operator
     * @param string $date
     * @return PathFilter
     */
    public function user_agent(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
MESSAGE* - текст сообщения которое будет выдано посетителю сайта, в случае его попадания под данную запись стоп-листа;

    /**
     * @param Operator $operator
     * @param string $date
     * @return PathFilter
     */
    public function message(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
COMMENTS* - административный комментарий, используется как правило для указания причин создания данной записи;

    /**
     * @param Operator $operator
     * @param string $date
     * @return PathFilter
     */
    public function comments(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
URL_TO* - страница (или ее часть) на которую приходит посетитель;

    /**
     * @param Operator $operator
     * @param string $date
     * @return PathFilter
     */
    public function url_to(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
URL_REDIRECT* - страница на которую необходимо перенаправить посетителя после его попадания под данную запись стоп-листа;

    /**
     * @param Operator $operator
     * @param string $date
     * @return PathFilter
     */
    public function url_redirect(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
SITE_ID* - ID сайта для которого запись будет действительна; если значение не задано, то это означает что запись действительная для всех сайтов;

    /**
     * @param Operator $operator
     * @param string $date
     * @return PathFilter
     */
public function site_id(Operator $operator, string $date): PathFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }

}