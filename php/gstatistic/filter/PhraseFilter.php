<?php


class PhraseFilter extends BaseFilter
{


    /**
     * UUID записи;
     *
     * @param Operator $operator
     * @param string $value
     * @return PhraseFilter
     */
    public function uuid(Operator $operator, string $value): PhraseFilter
    {
        $this->setFilter($operator, $value, 'uuid');
        return $this;
    }

    /**
     * UUID сессии;
     *
     * @param Operator $operator
     * @param string $value
     * @return PhraseFilter
     */
    public function sessionUuid(Operator $operator, string $value): PhraseFilter
    {
        $this->setFilter($operator, $value, 'sessionUuid');
        return $this;
    }

    /**
     * UUID поисковой системы;
     *
     * @param Operator $operator
     * @param string $value
     * @return PhraseFilter
     */
    public function searcherUuid(Operator $operator, string $value): PhraseFilter
    {
        $this->setFilter($operator, $value, 'searcherUuid');
        return $this;
    }

    /**
     * UUID записи из таблицы ссылающихся сайтов (страниц);
     *
     * @param Operator $operator
     * @param string $value
     * @return PhraseFilter
     */
    public function refererUuid(Operator $operator, string $value): PhraseFilter
    {
        $this->setFilter($operator, $value, 'refererUuid');
        return $this;
    }

    /**
     * Название поисковой системы;
     *
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    public function searcher(Operator $operator, string $value): PhraseFilter
    {
        $this->setFilter($operator, $value, 'searcher');
        return $this;
    }

    /**
     * Значение интервала для поля "дата"
     *
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    public function date(Operator $operator, string $value): PhraseFilter
    {
        $this->setFilter($operator, $value, 'date');
        return $this;
    }

    /**
     * Поисковая фраза;
     *
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    public function phrase(Operator $operator, string $value): PhraseFilter
    {
        $this->setFilter($operator, $value, 'phrase');
        return $this;
    }

    /**
     * Страница на которую пришли;
     *
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    public function to(Operator $operator, string $value): PhraseFilter
    {
        $this->setFilter($operator, $value, 'to');
        return $this;
    }

    /**
     * Была ли 404 ошибка на странице на которую пришли
     *
     * @param bool $value
     * @return $this
     */
    public function to404(bool $value): PhraseFilter
    {
        $this->setFilter(Operator::Eq, $value, 'to404');
        return $this;
    }

    /**
     * ID сайта, на который пришли;
     *
     * @param Operator $operator
     * @param string $value
     * @return $this
     */
    public function siteId(Operator $operator, string $value): PhraseFilter
    {
        $this->setFilter($operator, $value, 'siteId');
        return $this;
    }
}
//GROUP - группировка результирующего списка, возможные значения:
//P - группировка по поисковой фразе;
//S - группировка по поисковой системе.
//    public function group(Operator $operator, string $date): PhraseFilter
//    {
//        $this->setFilter($operator, $date, 'date');
//        return $this;
//    }