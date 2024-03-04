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
	fmt.Println("Pull data by", x.Host, d)
	for i := range d {
		x.Pull(&d[i])
	}

	return 1, nil
}

func Store(p *Piller, d []Data) (int, error) {
	fmt.Println("Store data by ", p.Host, d)
	for i := range d {
		p.Store(&d[i])
	}
	return 1, nil

}

type System struct {
	Xenia
	Piller
}

func copy(sys *System, batch int) {
	data := make([]Data, batch)

	Pull(&sys.Xenia, data)
	Store(&sys.Piller, data)
}

func main() {
	sys := System{Xenia: Xenia{Host: "local_Xenia", TimeOut: time.Second},
		Piller: Piller{Host: "local2_Piller", TimeOut: time.Second}}

	copy(&sys, 3)
}
