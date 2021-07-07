package goscheduler

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLockFreeSchedulerNormalSchedules(t *testing.T) {
	ch := make(chan int)
	lockTask := &mockLockTask{m: ch}
	lfs := NewLockFreeScheduler(lockTask, time.Millisecond*100)
	lfs.Start()
	runSuccess := waitForTask(2 * time.Second, 5, ch)
	assert.True(t, runSuccess, "lock scheduler run failed")

	lfs.Stop()
	didRun := waitForTask(1 * time.Second, 1, ch)
	assert.False(t, didRun, "after scheduler stop, should not run")
}

func waitForTask(waitTime time.Duration, runCnt int, ch <-chan int) bool {
	for {
		select {
		case <-time.After(waitTime):
			return false
		case <-ch:
			runCnt--
			if runCnt == 0 {
				return true
			}
		}
	}

}

type mockLockTask struct {
	m chan int
}

func (t *mockLockTask) Run() {
	t.m <- 0
}
