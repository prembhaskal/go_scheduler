package goscheduler

// Task is the job to be scheduled
type Task interface {
	Run()
}

// SchedulerIntf interface for all different kind of schedulers
type SchedulerIntf interface {
	Start()
	Stop()
}
