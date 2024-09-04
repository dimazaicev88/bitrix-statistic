<?php


class GAdv
{

    /**
     * Удаление рекламной компании
     *
     * @param string $uuid
     * @return void
     */
    public static function delete(string $uuid): void
    {

    }

    /**
     * @param AdvFilter $filter
     * @param array $fields
     * @param array $order
     * @param string $orderBy
     * @param int $skip
     * @param int $limit
     * @return array
     */
    public static function find(AdvFilter $filter, array $fields = [], array $order = [], string $orderBy = "", int $skip = 0, int $limit = 0): array
    {
        return [];
    }

    /**
     * @param string $uuid
     * @return array
     */
    public static function findByUuid(string $uuid): array
    {
        return [];
    }

    /**
     * Возвращает список типов событий инициализированным посетителями зашедшими по определённой рекламной кампании.
     *
     * @param string $advUuid
     * @param string $by
     * @param string $order
     * @param array $arFilter
     * @return array
     */
    public static function getEventList(string $advUuid, string $by = 's_counter', string $order = 'desc', array $arFilter = []): array
    {
        return [];
    }

    public static function GetEventListByReferer($value, $arFilter): array
    {
        return [];
    }

    /**
     * Возвращает данные по трафику рекламной кампании в разрезе по дням.
     *
     * @param string $advUuid
     * @param string $by
     * @param string $order
     * @param array $arMaxMin
     * @param array $arFilter
     * @return array
     */
    public static function GetDynamicList(string $advUuid, string $by = 's_date', string $order = 'desc', array &$arMaxMin = [], array $arFilter = []): array
    {
        return [];
    }

    public static function GetDropDownList($strSqlOrder = "ORDER BY REFERER1, REFERER2"): array
    {
        return [];
    }

    /**
     * Возвращает список рекламных кампаний.
     *
     * @param string $by
     * @param string $order
     * @param array $arFilter
     * @return array
     */
    public static function GetSimpleList(string $by = 's_referer1', string $order = 'asc', array $arFilter = []): array
    {
        return [];
    }
}