package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()

	var t *time.Timer
	ch := make(chan bool)
	t = time.AfterFunc(randomDur(), func() {
		fmt.Println(time.Now().Sub(start))
		ch <- true
	})
	//for time.Since(start) < 10*time.Second {
	<-ch
	t.Reset(randomDur())
	//}

	time.Sleep(5 * time.Second)
}

func randomDur() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}
