// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package clock

import (
	"sync"
	"time"
)

type Clock interface {
	Now() time.Time
}

type RealClock struct{}

func (RealClock) Now() time.Time {
	return time.Now()
}

type FakeClock struct {
	mutex sync.RWMutex
	time  time.Time
}

func NewFakeClock(t time.Time) *FakeClock {
	return &FakeClock{
		time: t,
	}
}

func (f *FakeClock) Now() time.Time {
	f.mutex.RLock()
	defer f.mutex.RUnlock()

	return f.time
}

func (f *FakeClock) SetTime(t time.Time) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	f.time = t
}

func (f *FakeClock) Step(d time.Duration) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	f.time = f.time.Add(d)
}
