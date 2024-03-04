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
type Piller struct {
	Host    string
	TimeOut time.Duration
}

func (Xenia) Pull(d *Data) error {
	fmt.Println("Pulling data", d.Line)
	return nil
}

func (Piller) Store(d *Data) error {
	fmt.Println("Storing data", d.Line)
	return nil
}

func Pull(x *Xenia, d []Data) (int, error) {
	fmt.Println("Pull data", d)
	for i := range d {
		x.Pull(&d[i])
	}

	return 1, nil
}

func Store(p *Piller, d []Data) (int, error) {
	fmt.Println("Store data", d)
	for i := range d {
		p.Store(&d[i])
	}
	return 1, nil

}
