package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {

	retryTimeout(context.Background(), time.Minute, func(ctx context.Context) error {
		return errors.New("error ")
	})
}

func retryTimeout(ctx context.Context, retryInterval time.Duration, check func(ctx context.Context) error) {
	for {
		fmt.Println("perform user check call")
		if err := check(ctx); err == nil {
			fmt.Println("work finished successfully")
			return
		}
		fmt.Println("check if timeout has expired")
		if ctx.Err() != nil {
			fmt.Println("time expired 1 :", ctx.Err())
			return
		}
		fmt.Printf("wait %s before trying again\n", retryInterval)
		t := time.NewTimer(retryInterval)
		select {
		case <-ctx.Done():
			fmt.Println("timed expired 2 :", ctx.Err())
			t.Stop()
			return
		case <-t.C:
			fmt.Println("retry again")
		}
	}
}
