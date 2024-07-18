package cache

import (
	"bitrix-statistic/internal/config"
	"sync"
)

var (
	mutex       sync.Mutex
	optionCache [1]config.AppOptions
)

func SetOptions(options config.AppOptions) {
	mutex.Lock()
	defer mutex.Unlock()
	optionCache[0] = options
}

func GetOptions() config.AppOptions {
	mutex.Lock()
	defer mutex.Unlock()
	return optionCache[0]
}
