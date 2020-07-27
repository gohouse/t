package t

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var defaultRecover = func(err error) {
	log.Println("panic err catch:", err)
}

func SetDefaultRecover(dr func(err error)) {
	defaultRecover = dr
}
func WithRecover(h func(), errDefaultRecover ...func(err error)) {
	defer func() {
		if err := recover(); err != nil {
			if len(errDefaultRecover) > 0 {
				errDefaultRecover[0](fmt.Errorf("%v", err))
			} else {
				defaultRecover(fmt.Errorf("%v", err))
			}
		}
	}()
	h()
}

var mu *sync.Mutex
var mur *sync.RWMutex

func WithLockContext(ctx func()) {
	mu.Lock()
	defer mu.Unlock()
	ctx()
}
func WithRLockContext(ctx func()) {
	mur.RLock()
	defer mur.RUnlock()
	ctx()
}

func WithTicker(duration time.Duration, fn func()) {
	ticker := time.NewTicker(duration)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			fn()
		}
	}
}
