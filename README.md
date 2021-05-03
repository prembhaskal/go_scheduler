# go_scheduler
yet another scheduler library in go

## Scheduler
type `goscheduler.Scheduler` can schedule a given `goscheduler.Task` periodically. In addition scheduler can be restarted once it is stopped, which is useful in some scenarios.  
API is as follows:  
`NewScheduler()` to create the scheduler instance  
`Start()` to start the scheduler  
`Stop()` to stop the scheduler  

### test
run the test using make command.
```
make test
```