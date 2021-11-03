package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main1() {
	start := time.Now()

	var t *time.Timer
	t = time.AfterFunc(randomDur1(), func() {
		fmt.Println(time.Now().Sub(start))
		t.Reset(randomDur1())
	})
	time.Sleep(5 * time.Second)
}

func randomDur1() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}
