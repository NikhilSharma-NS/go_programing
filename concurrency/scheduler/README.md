#### M:N Scheduler

The Go scheduler is part of the Go runtime. it is known as M:N scheduler
Go Scheduler runs in user space
Go Scheduler uses OS threads to schedule goroutines for execution 
Go routines runs in the context of os thread
Go runtime create number of worker OS thraeds equal to GOMAXPROCS
GOMAXPROCS - default value is number of processors on machine
Go scheduler distributes runnable goroutines over multiple worker OS threads
At any time , N goroutines could be scheduled on M OS threads that runs on at most GOMAXPROCS number of processors.

#### asynchronous preemption
- As of Go 1.14 Go Scheduler implements asynchronous preemption
- this prevents long running goroutines from hogging onto CPU,that could block other Goroutines
- the asynchronous preemption is triggered based on time condition. When a goroutine is running for more than 10ms Go will try to preempt it


```
    Core
     |
     M ------- G1  local run queue
     |
     P ------- G2                              G10                 Global rune queue
                                               G11


M : OS Thread
P : Processor
G : goroutine

```

#### Context Switch due to Synchronous System Call

a1) what happens in general when sync system call are made

- sync system calls wait for I/O operation to be complted
- OS thread is moved out of the CPU to waiting queue for I/O to complte
- sync system call reduces parallelism

b1) what happens in general when async system call are made

- File descriptor is set to non-blocking mode
- If file descriptor is not ready for I/O operation , system call does not block but return an error
- async IO increses the application complexity
- Setup event loops using callbacks functions

- Go uses netpoller to handle async system call
- netpoller uses interface provided by OS to do polling on file descriptor and notifies the gorountine to try I/O operation when it ready
- App complexity of managing async system call is moved to Go runtime which manages it efficiently

##### Work Stealing

If there is no goroutines in local run queue then 
1) try to steal from other logical processor
2) if not found check the global runtime queue for a G
3) if not found check netpoller


