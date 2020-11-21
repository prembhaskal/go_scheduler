package scheduler

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSimpleSchedule(t *testing.T) {
	t.Log("started simple schedule test.")
	ch := make(chan int)
	mockTask := &mockTask{
		ch: ch,
	}

	interval := 100 * time.Millisecond
	runCount := 5

	scheduler := NewScheduler(mockTask, interval)
	scheduler.Start()
	assert.True(t, scheduler.started, "scheduler did not started")

	runStatus := waitForTaskToRun(time.Duration(2*runCount)*interval, runCount, ch)
	assert.True(t, runStatus, "scheduler did not run properly")

	scheduler.Stop()
	assert.False(t, scheduler.started, "scheduler did not stop")
}

func waitForTaskToRun(waitTime time.Duration, runCount int, waitCh <-chan int) bool {
	for {
		select {
		case <-time.After(waitTime):
			return false
		case <-waitCh:
			runCount--
			if runCount == 0 {
				return true
			}
		}
	}
}

type mockTask struct {
	ch chan int
}

func (m *mockTask) Run() {
	m.ch <- 0
}
