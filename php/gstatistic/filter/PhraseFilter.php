<?php


class PhraseFilter extends BaseFilter
{


    /**
     * UUID записи;
     *
     * @param Operator $operator
     * @param string $uuid
     * @return PhraseFilter
     */
    public function uuid(Operator $operator, string $uuid): PhraseFilter
    {
        $this->setFilter($operator, $uuid, 'uuid');
        return $this;
    }

    /**
     * UUID сессии;
     *
     * @param Operator $operator
     * @param string $sessionUuid
     * @return PhraseFilter
     */
    public function sessionUuid(Operator $operator, string $sessionUuid): PhraseFilter
    {
        $this->setFilter($operator, $sessionUuid, 'sessionUuid');
        return $this;
    }

    /**
     * UUID поисковой системы;
     *
     * @param Operator $operator
     * @param string $searcherUuid
     * @return PhraseFilter
     */
    public function searcherUuid(Operator $operator, string $searcherUuid): PhraseFilter
    {
        $this->setFilter($operator, $searcherUuid, 'searcherUuid');
        return $this;
    }

    /**
     * UUID записи из таблицы ссылающихся сайтов (страниц);
     *
     * @param Operator $operator
     * @param string $refererUuid
     * @return PhraseFilter
     */
    public function refererUuid(Operator $operator, string $refererUuid): PhraseFilter
    {
        $this->setFilter($operator, $refererUuid, 'refererUuid');
        return $this;
    }

    /**
     * Название поисковой системы;
     *
     * @param Operator $operator
     * @param string $date
     * @return $this
     */
    public function searcher(Operator $operator, string $date): PhraseFilter
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

    /**
     * Поисковая фраза;
     *
     * @param Operator $operator
     * @param string $date
     * @return $this
     */
    public function phrase(Operator $operator, string $date): PhraseFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }

    /**
     * Страница на которую пришли;
     *
     * @param Operator $operator
     * @param string $date
     * @return $this
     */
    public function to(Operator $operator, string $date): PhraseFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }

    /**
     * Была ли 404 ошибка на странице на которую пришли
     *
     * @param bool $date
     * @return $this
     */
    public function to404(bool $date): PhraseFilter
    {
        $this->setFilter(Operator::Eq, $date, 'date');
        return $this;
    }

    /**
     * ID сайта, на который пришли;
     *
     * @param Operator $operator
     * @param string $date
     * @return $this
     */
    public function siteId(Operator $operator, string $date): PhraseFilter
    {
        $this->setFilter($operator, $date, 'date');
        return $this;
    }
}
//GROUP - группировка результирующего списка:
//P - группировка по поисковой фразе;
//S - группировка по поисковой системе.
//    public function group(Operator $operator, string $date): PhraseFilter
//    {
//        $this->setFilter($operator, $date, 'date');
//        return $this;
//    }