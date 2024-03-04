package main

import (
	"fmt"
	"time"
)

type Data struct {
	Line string
}
type Xenia struct {
	Host    string
	TimeOut time.Duration
}

func (Xenia) Pull(d *Data) error {
	fmt.Println("Pulling data", d.Line)
	return nil
}

type Piller struct {
	Host    string
	TimeOut time.Duration
}

func (Piller) Store(d *Data) error {
	fmt.Println("Storing data", d.Line)
	return nil
}
