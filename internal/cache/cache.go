package cache

import (
	"github.com/maypok86/otter"
	"github.com/sirupsen/logrus"
	"sync"
)

var (
	once       sync.Once
	otterCache otter.Cache[string, interface{}]
)

func Cache() otter.Cache[string, interface{}] {
	once.Do(func() {
		var err error
		otterCache, err = otter.MustBuilder[string, interface{}](100).CollectStats().Build()

		if err != nil {
			logrus.Fatal(err)
		}
	})

	return otterCache
}

func Close() {
	otterCache.Close()
}
