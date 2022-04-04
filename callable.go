package gc

import (
	"sync"
	"time"
)

var (
	callableInit  bool
	callableItems map[string]Callable
	callableMutex sync.RWMutex
)

type Callable func() (v interface{}, expAt *time.Time, err error)

// SetCallable allows defining callable function for giving gc's name.
// When a value expire, if one callback is defined, gc call it for get a fresh value.
func SetCallable(gcName string, c Callable) {
	callableMutex.Lock()

	initCallable()

	callableItems[gcName] = c

	callableMutex.Unlock()
}

func initCallable() {
	if !callableInit {
		callableItems = make(map[string]Callable)
		callableInit = true
	}
}
