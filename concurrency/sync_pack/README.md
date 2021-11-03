#### Sync Package


when to use channels and when to use mutex

Channels:
a) passing copy to data
b) Dirtributing units to work
c) Communicating asynchronous results


Mutex:
a) Cache
b) State

- Mutex used for protect shared resources
- sync.Mutex : Provide exclusive access to a shared resource

```
mu.Lock()
balance+=amount
mu.Unlock()

```

```

mu.Lock()
defer mu.Unlock()
balance-=amount

```

- sync.RWMutex - allows multiple readers . Writers get exclusive lock.

```
mu.Lock()
balance+=amount
mu.Unlock()
```

```
mu.Rlock()
defer mu.RUnlock()
return balance
```

#### sync.Atomic

- low level atomic operations on memory
- Lockless operation
- Used for atomic operations on counters.

```
atomic.AddUint64(&ops,10)
value:=atomic.LoadUint64(&ops)
```

```
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	runtime.GOMAXPROCS(4)

	var counter uint64
	var wg sync.WaitGroup

	for couneter1 := 0; couneter1 < 50; couneter1++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&counter, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("counter", counter)

}

```

#### sync.Cond

- Condition variable is one of the synchronization mechanisms
- a condition variable is basically a container of Goroutines that are waiting for a certain condition

a) How to make a goroutine wait till some event(condition) occur?
a.1) Wait in loop for the condition
```
var sh=make(map[string]string)
go func(){
    defer wg.Done()
    mu.Lock()
    for len(sh)==0{
        mu.Unlock()
        time.Sleep(100*time.Millisecond)
        mu.Lock()
    }
    fmt.Println(sh["rsc"])
    mu.UnLock()
}()
```

- we need some way to make goroutine suspend while waiting
- we need some way to signal the suspended goroutine when particular event has occured.
Channels 
- we can use channels to block a goroutine on receive
- Sender go routine to indicate occurence of event
- what if there are multiple goroutines waiting on multiple conditions/event

sync.Cond
- Conditional variable are type
```
var c *sync.Cond
```
- we can use constrctor method sync.NewCond() to create a conditional variable
- it takes sync.Locker interface as input, which is usually sync.Mutex

```
m:= sync.Mutex{}
c:= sync.NewCond(&m)
```
- it has 3 method
```
c.Wait()
c.Signal()
c.Broadcast()
```
- c.Wait()
  - suspends execution of the calling goroutine
  - automatically unlocks c.L
  - wait cannot return unless awoken by Broadcast or signal
  - Wait locks c.L before returing
  - Because c.L is not locked when wait first resumes the caller typically cannot assume that condition is true
  when wait returns . instead the caller should wait in a loop

- c.Signal()
  - singal wakes one goroutine waiting on c if there is any
  - signal finds goroutine that has been waiting the longest and notifies that
  - it is allowed but not required for the caller to hold c.L during the call

- c.BroadCast()
  - broadcast wakes all goroutines waiting on c
  - it is allowed but not required for the caller to hold c.L during the call


#### sync.Once

- Run one-time initialization functions
- once.Do(funcValue)
- sync.Once ensure that only one call to Do ever calls the function passed in - even 
on different goroutines

```
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var synco sync.Once

	load := func() {
		fmt.Println("Run Once")
	}
	loadmul := func() {
		fmt.Println("Run multiple")
	}

	for counter := 0; counter < 10; counter++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			synco.Do(load)
			loadmul()
		}()
	}
	wg.Wait()

}

```

#### sync.Pool

- create and make avaliable pool of things for use

```
b:=bufPool.Get().(*bytes.Buffer)
```

```
bufPool.Put(b)
```
