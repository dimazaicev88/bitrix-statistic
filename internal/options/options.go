package options

import (
	"bitrix-statistic/internal/cache"
	"github.com/sirupsen/logrus"
)

func Set(key string, value interface{}) {

	switch value.(type) {

	case int:
	case float64:
	case string:
	default:
		logrus.Panic("unknown type")
	}

	cache.Cache().Set(key, value)
}
