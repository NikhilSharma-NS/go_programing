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

type Puller interface {
	Pull(d *Data) error
}

type Storer interface {
	Store(d *Data) error
}

func (Xenia) Pull(d *Data) error {
	fmt.Println("Pulling data", d.Line)
	return nil
}

func (Piller) Store(d *Data) error {
	fmt.Println("Storing data", d.Line)
	return nil
}

func Pull(p Puller, d []Data) (int, error) {
	fmt.Println("Pull data ", d)
	for i := range d {
		p.Pull(&d[i])
	}

	return 1, nil
}

func Store(s Storer, d []Data) (int, error) {
	fmt.Println("Store data  ", d)
	for i := range d {
		s.Store(&d[i])
	}
	return 1, nil

}

func copy(p Puller, s Storer, batch int) {
	data := make([]Data, batch)

	Pull(p, data)
	Store(s, data)
}
func main() {
	x := Xenia{
		Host:    "local_Xenia",
		TimeOut: time.Second,
	}
	p := Piller{
		Host:    "local2_Piller",
		TimeOut: time.Second,
	}

	copy(x, p, 3)
}