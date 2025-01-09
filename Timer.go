package src

import "time"

type Timer struct {
	Duration time.Duration
	Handler  func()
}

func (t *Timer) Start() {
	go func() {
		time.Sleep(t.Duration)
		t.Handler()
	}()
}
