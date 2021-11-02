#### Select

G1 wants to receive result of computation from G2 and G3

```
        G2   Task1
G1
        G3   Task2

```

in what order are we going to receive results

g2 or g3
g3 or g2

can we do operation on channel which ever is raedy and don't worry about the order?

select is like a switch statement
Each cases specifies communication
all channel opeartion are considered simultaneously.
select wait until some case is raedy to processed.
when one channels is raedy that operation will proceed
select is also very helpful in implementing
  - Timout 
  - Non Blocking communication

```

select {
    case <-ch1:
    // block of statement
    case <-ch2:
    // block of statement
    case ch3 <- struct{}{}:
   // block of statement

}
```

```
select{
    case v: <-ch:
     fmt.Println(v)
    case <-time.After(3*time.Second):
     fmt.Println("timeout")
}

```

- select waits untill there is event on channel ch or until timeout is reached

- The time.After function take in a time. Duration argument and returns a channel that will send 
the current time after the duration we provide it

##### Non-blocking communication
```
select {
    case m:= <-ch:
      fmt.Println("received",m)
    default:
      fmt.Println("no msg received")
}
```

send or receive on a channel but avoid blocking if the channel is not raedy

