package gc

import (
	"sync"
	"time"
)

type cache struct {
	Value     interface{}
	CreatedAt int64
	ExpireAt  int64
}

var (
	cacheInit  bool
	cacheItems map[string]cache
	cacheMutex sync.RWMutex
)

func Set(name string, value interface{}, expireAt *time.Time) {
	cacheMutex.Lock()

	initGc()

	cacheItems[name] = cache{
		Value:     value,
		CreatedAt: time.Now().UnixNano(),
		ExpireAt:  expireAtToUnix(expireAt),
	}

	cacheMutex.Unlock()
}

func Get(name string) interface{} {
	cacheMutex.RLock()

	initGc()

	if cacheItems[name].ExpireAt != 0 && cacheItems[name].ExpireAt <= time.Now().UnixNano() {
		cacheMutex.RUnlock()
		cacheMutex.Lock()
		delete(cacheItems, name)
		cacheMutex.Unlock()
		cacheMutex.RLock()
	}

	val := cacheItems[name].Value

	cacheMutex.RUnlock()

	if val == nil && callableItems[name] != nil {
		v, expAt, err := callableItems[name]()
		if err == nil {
			Set(name, v, expAt)
			return v
		}
	}

	return val
}

func initGc() {
	if !cacheInit {
		cacheItems = make(map[string]cache)
		cacheInit = true
	}
}

func expireAtToUnix(expireAt *time.Time) int64 {
	if expireAt == nil {
		return 0
	} else {
		return expireAt.UnixNano()
	}
}
