package main

import (
	"context"
	"math/rand"

	"fmt"
	"time"
)

type workfn[Result any] func(context.Context) Result

func doWork[Result any](ctx context.Context, work workfn[Result]) chan Result {
	ch := make(chan Result, 1)

	go func() {
		ch <- work(ctx)
		fmt.Println("do: done")
	}()
	return ch
}
func main() {
	d := 100 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), d)
	defer cancel()

	fwf := func(ctx context.Context) string {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return "done"
	}

	res := doWork(ctx, fwf)

	select {
	case v := <-res:
		fmt.Println("main", v)
	case <-ctx.Done():
		fmt.Println("timeout")
	}

}
