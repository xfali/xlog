// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package timer

import (
	"sync"
	"time"
)

type RecordTimer struct {
	current time.Time
	lock    sync.RWMutex

	stop chan struct{}
	once sync.Once
}

func NewRecordTimer(interval time.Duration) *RecordTimer {
	ret := &RecordTimer{
		stop: make(chan struct{}),
	}
	go func() {
		tick := time.NewTicker(interval)
		defer tick.Stop()
		for {
			select {
			case <-tick.C:
				ret.refresh()
			case <-ret.stop:
				return
			}
		}
	}()

	return ret
}

func (t *RecordTimer) refresh() {
	t.lock.Lock()
	t.current = time.Now()
	t.lock.Unlock()
}

func (t *RecordTimer) Now() time.Time {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.current
}

func (t *RecordTimer) Stop() bool {
	t.once.Do(func() {
		close(t.stop)
	})
	return false
}
