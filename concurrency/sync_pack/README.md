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
defer mu.Runlock()
return balance
```