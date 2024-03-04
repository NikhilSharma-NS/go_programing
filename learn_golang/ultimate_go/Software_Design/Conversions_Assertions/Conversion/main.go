package main

import "fmt"

type Mover interface {
	Move()
}

type Locker interface {
	Lock()
	Unlock()
}

type MoveLocker interface {
	Mover
	Locker
}

type bike struct{}

func (bike) Move() {
	fmt.Println("bike moving")
}

func (bike) Lock() {
	fmt.Println("bike locking")
}

func (bike) Unlock() {
	fmt.Println("bike Unlocking")
}

func main() {
	var m1 MoveLocker
	var m Mover

	m1 = bike{}
	m = m1

	//m1 = m
	// cannot use m (variable of type Mover) as MoveLocker value in assignment: missing method LockcompilerInvalidIfaceAssign

	fmt.Println(m)

}
