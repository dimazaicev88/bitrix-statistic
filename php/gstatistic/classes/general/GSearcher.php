<?php

class GSearcher
{
    /**
     * @param string $searcherUuid
     * @param string $date1
     * @param string $date2
     * @return array
     */
    public static function dynamicDays(string $searcherUuid, string $date1 = "", string $date2 = ""): array
    {
        return [];
    }

    public static function getGraphArray($arFilter, &$arrLegend): array
    {
        return [];
    }

    /**
     * Возвращает список доменов поисковых систем.
     *
     * @param string $by
     * @param string $order
     * @param array $arFilter
     * @return array
     */
    public static function getDomainList(string $by = 's_id', string $order = 'desc', array $arFilter = []): array
    {
        return [];
    }

    /**
     * Возвращает данные по указанной поисковой системе.
     *
     * @param string $uuid
     * @return void
     */
    public static function getByUuid(string $uuid)
    {
    }

}