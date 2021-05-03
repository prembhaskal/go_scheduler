package goscheduler

import (
	"time"
)

// LockFreeScheduler is scheduler implementation which does not used mutex for shared locking
type LockFreeScheduler struct {
	interval time.Duration
	chListen chan bool
	chStop   chan bool
	task     Task
}

// NewLockFreeScheduler creates new instance of LockFreeScheduler
func NewLockFreeScheduler(task Task, interval time.Duration) *LockFreeScheduler {
	s := &LockFreeScheduler{
		interval: interval,
		chListen: make(chan bool, 1),
		chStop:   make(chan bool, 1),
		task:     task,
	}
	s.startRoutine()
	return s
}

// Start starts the task scheduling, it does nothing if scheduler is already started.
func (s *LockFreeScheduler) Start() {
	s.chListen <- true
}

// Stop stops the scheduler. it does nothing if scheduler is already stopped.
func (s *LockFreeScheduler) Stop() {
	s.chListen <- false
}

func (s *LockFreeScheduler) startRoutine() {
	var isStarted bool
	for {
		state := <-s.chListen // reads one msg at a time, blocks any further reads unless a proper state is established
		if state {
			if !isStarted {
				// post to start polling
				go s.poll()
				isStarted = true
			}
		} else {
			if isStarted {
				// post to stop polling
				s.chStop <- true
				isStarted = false
			}
		}
	}
}

func (s *LockFreeScheduler) poll() {
	ticker := time.NewTicker(s.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			s.task.Run()
		case <-s.chStop:
			return
		}
	}
}
