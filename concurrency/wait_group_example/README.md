### Wait Group

```
main

 -> go -> go rountine  -> 

 join point <-
```


Step 1: Declare the waitGroup
```
	var wg sync.WaitGroup

```
Step 2: Add the integer value using Add method 

```
wg.Add(1)
```

Step 3: in main function or main routine we need to add wait statement

```
	wg.Wait()

```

Step 4 : we have to defer statement in go routine to consider or complete the go routine

```
		defer wg.Done()

```

Step 5 : final program

```
package main

import (
	"fmt"
	"sync"
)

func main() {

	var data int
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		data++
	}()

	wg.Wait()
	fmt.Printf("data value is %v", data)

}

```