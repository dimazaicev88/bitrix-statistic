<?php

class GPage
{

    /**
     * Возвращает данные по посещаемости указанной страницы в разрезе по дням.
     * @return array
     */
    public static function getDynamicList(): array
    {
        return [];
    }

    public static function find(PageFilter $filter, array $fields = [], array $order = [], string $orderBy = "", int $skip = 0, int $limit = 0): array
    {
        return [];
    }

}