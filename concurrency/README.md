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


