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

func WithLockContext(mu *sync.Mutex, fn func()) {
	mu.Lock()
	defer mu.Unlock()
	fn()
}
func WithRLockContext(mur *sync.RWMutex, fn func()) {
	mur.RLock()
	defer mur.RUnlock()
	fn()
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

func WithRunTimeContext(closer func(), callback func(time.Duration)) {
	// 记录开始时间
	start := time.Now()
	closer()
	timeduration := time.Since(start)
	//log.Println("执行完毕,用时:", timeduration.Seconds(),timeduration.Seconds()>1.1)
	callback(timeduration)
}
