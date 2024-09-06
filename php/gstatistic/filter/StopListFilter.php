<?php

class StopListFilter extends BaseFilter
{

    /**
     * UUID записи стоп-листа
     *
     * @param Operator $operator
     * @param string $value
     * @return StopListFilter
     */
    public function uuid(Operator $operator, string $value): StopListFilter
    {
        $this->setFilter($operator, $value, 'uuid');
        return $this;
    }

    /**
     * Время начала активности записи
     *
     * @param Operator $operator
     * @param string $value
     * @return StopListFilter
     */
    public function dateStart(Operator $operator, string $value): StopListFilter
    {
        $this->setFilter($operator, $value, 'dateStart');
        return $this;
    }

    /**
     * Время окончания активности записи
     *
     * @param Operator $operator
     * @param string $value
     * @return StopListFilter
     */
    public function dateEnd(Operator $operator, string $value): StopListFilter
    {
        $this->setFilter($operator, $value, 'dateEnd');
        return $this;
    }


    /**
     * Флаг активности записи
     *
     * @param bool $value
     * @return StopListFilter
     */
    public function active(bool $value): StopListFilter
    {
        $this->setFilter(Operator::Eq, $value, 'active');
        return $this;
    }

    /**
     * Флаг необходимости сохранения статистики по посетителю попавшему в стоп-лист, воможные значения
     *
     * @param bool $value
     * @return StopListFilter
     */
    public function saveStatistic(bool $value): StopListFilter
    {
        $this->setFilter(Operator::Eq, $value, 'saveStatistic');
        return $this;
    }

    /**
     * Октет 1 IP адреса
     *
     * @param Operator $operator
     * @param string $value
     * @return StopListFilter
     */
    public function ip1(Operator $operator, string $value): StopListFilter
    {
        $this->setFilter($operator, $value, 'ip1');
        return $this;
    }

    /**
     * Октет 2 IP адреса
     *
     * @param Operator $operator
     * @param string $value
     * @return StopListFilter
     */
    public function ip2(Operator $operator, string $value): StopListFilter
    {
        $this->setFilter($operator, $value, 'ip2');
        return $this;
    }

    /**
     * Октет 3 IP адреса
     *
     * @param Operator $operator
     * @param string $value
     * @return StopListFilter
     */
    public function ip3(Operator $operator, string $value): StopListFilter
    {
        $this->setFilter($operator, $value, 'ip3');
        return $this;
    }

    /**
     * Октет 4 IP адреса
     *
     * @param Operator $operator
     * @param string $value
     * @return StopListFilter
     */
    public function ip4(Operator $operator, string $value): StopListFilter
    {
        $this->setFilter($operator, $value, 'ip4');
        return $this;
    }

    /**
     * Ссылающаяся страница, с которой приходит посетитель
     *
     * @param Operator $operator
     * @param string $value
     * @return StopListFilter
     */
    public function urlFrom(Operator $operator, string $value): StopListFilter
    {
        $this->setFilter($operator, $value, 'urlFrom');
        return $this;
    }


    /**
     * UserAgent посетителя
     *
     * @param Operator $operator
     * @param string $value
     * @return StopListFilter
     */
    public function userAgent(Operator $operator, string $value): StopListFilter
    {
        $this->setFilter($operator, $value, 'userAgent');
        return $this;
    }


    /**
     * Текст сообщения которое будет выдано посетителю сайта, в случае его попадания под данную запись стоп-листа
     *
     * @param Operator $operator
     * @param string $value
     * @return StopListFilter
     */
    public function message(Operator $operator, string $value): StopListFilter
    {
        $this->setFilter($operator, $value, 'message');
        return $this;
    }

    /**
     * Административный комментарий, используется как правило для указания причин создания данной записи
     *
     * @param Operator $operator
     * @param string $value
     * @return StopListFilter
     */
    public function comments(Operator $operator, string $value): StopListFilter
    {
        $this->setFilter($operator, $value, 'comments');
        return $this;
    }


    /**
     * Страница (или ее часть) на которую приходит посетитель
     *
     * @param Operator $operator
     * @param string $value
     * @return StopListFilter
     */
    public function urlTo(Operator $operator, string $value): StopListFilter
    {
        $this->setFilter($operator, $value, 'urlTo');
        return $this;
    }


    /**
     * Страница на которую необходимо перенаправить посетителя после его попадания под данную запись стоп-листа
     *
     * @param Operator $operator
     * @param string $value
     * @return StopListFilter
     */
    public function urlRedirect(Operator $operator, string $value): StopListFilter
    {
        $this->setFilter($operator, $value, 'urlRedirect');
        return $this;
    }

    /**
     * ID сайта для которого запись будет действительна, если значение не задано, то это означает что запись действительная для всех сайтов
     *
     * @param Operator $operator
     * @param string $value
     * @return StopListFilter
     */
    public function siteId(Operator $operator, string $value): StopListFilter
    {
        $this->setFilter($operator, $value, 'siteId');
        return $this;
    }

}