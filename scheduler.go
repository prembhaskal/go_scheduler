package scheduler

import (
	"sync"
	"time"
)

// Task is the job to be scheduled
type Task interface {
	Run()
}

// Scheduler schedules the task periodically. Scheduler can be restarted once stopped.
type Scheduler struct {
	pollingInterval time.Duration
	chStop          chan bool
	mutex           sync.Mutex
	started         bool
	task            Task
}

// Start starts the task scheduling, it does nothing if scheduler is already started.
func (s *Scheduler) Start() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if !s.started {
		go s.poll()
		s.started = true
	}
}

// Stop stops the scheduler. it does nothing if scheduler is already stopped.
func (s *Scheduler) Stop() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.started {
		s.chStop <- true
		s.started = false
	}
}

func (s *Scheduler) poll() {
	ticker := time.NewTicker(s.pollingInterval)
	defer ticker.Stop()

	for {
		select {
		case <- ticker.C:
			s.task.Run()
		case <- s.chStop:
			return
		}
	}
}
