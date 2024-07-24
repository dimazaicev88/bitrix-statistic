package entityjson

import "github.com/google/uuid"

type Path struct {
	PathId         uuid.UUID // ID отрезка пути
	LastPage       string    // Последняя страница отрезка пути
	LastPage404    bool      // True|False флаг 404 ошибки на последней странице пути
	LastPageSiteId string    // ID сайта последней страницы пути
	Counter        uint32    // Количество переходов по отрезку пути
}

type PathFull struct {
	PathId  uuid.UUID // ID отрезка пути
	Pages   string    // Набор страниц входящих в полный путь
	Counter uint32    // Количество переходов по полному пути
}
