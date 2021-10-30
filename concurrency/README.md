# Introduction

Concurrency is composition of independent execution computations,which may or may not run in parallel

Parallelism is ability to execute multiple computations 


# Process and Thread

Opearting System
the job of os system is to give fair chance for all processes access to CPU,memory and other resources

Process
An instance of a running program is called a process

process provides environment for program to execute 

OS allocates memory 
Code - machine instruction
Data - Global data
Heap - Dynamic memory allocation
Stack - Local variable of function

Thread

threads are smallest unit of execution that CPU accepts
process has atleast one thread -> main thread
process can have multiple threads
Threads share same address space
Threads run independent of each other
OS schedular makes scheduling decisions at thraed level not process level
thraeds can run concurrently or parallel

```
        <- preempted <-

Runnable -> Executing  

          <-  waiting <-

```
C10K Problems

Scheduler allocates a process a time slice for execution on CPU core

this CPU time slice is divided equally amon threads

```
Scheduler time    Number of thread   Thraed time slice
10 ms                    2               5ms
10 ms                    5               2ms
10 ms                    1000            10us

```

if minimum time for thread is slice 2ms

```
Scheduler time    Number of thread   Thraed time slice
2 s                    1000               2ms
20 s                   10000              2ms


```
Fixed Stack Size

Thraeds are allocated fixed stack size 

## Why Concurrency is Hard

Shared Memory

thread communication between each other by sharing memory
sharing of memory between threads creates lots of complexity 
concurrent access to shared memory by 2 or more thread can lead to Data Race and outcome can be Un-deterministic

# GoRountines


Communicating Sequenial Processes (SCP)

1) Each process is built for sequential Execution
2) Data is communicated between processes . No Shared Memory
3) Scale by adding more of the same

tool set in golang
1) goroutines
2) channels
3) select
4) sync pacakage

#### Goroutines
1) we can think gorountines as user space thread managed by go runtime
2) Gorountines extermely lightweight. Go routines starts with 2kb stack
which grows and shrinks as required
3) Low CPU overhead 3 instructions per function call
4) Can create 100 of 1000 of gorountines in the same address space
5) Channels are used for communications of data between goroutines 
Sharing of memory can be avoided
6) Context switching between goroutines is much cheaper than thread context switching
7) Go runtime can be more selective in what is persisted for retrieval
how it is persisted and when the persisting needs to occur
8) Go runtime creates worker OS threads 
9) Goroutines runs in the context of OS thread
10) Many goroutines executes in the context of single OS thread
