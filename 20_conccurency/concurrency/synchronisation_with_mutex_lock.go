package concurrency

import (
	"sync"
	"time"
)

func SynchronisationMutex() {
	s := &Services{}
	s.Start()
	time.Sleep(time.Second)
	s.Stop()
}

type Services struct {
	started bool
	stpCh   chan struct{}
	mutex   sync.Mutex
}

func (s *Services) Start() {
	s.stpCh = make(chan struct{})
	go func() {
		s.mutex.Lock()
		s.started = true
		s.mutex.Unlock()
		<-s.stpCh
	}()
}

func (s *Services) Stop() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.started {
		s.started = false
		close(s.stpCh)
	}
}
