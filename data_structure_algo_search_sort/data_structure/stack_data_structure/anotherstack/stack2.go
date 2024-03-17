package main

import "fmt"

type Stack2 struct {
	items []int
}


func (s *Stack2) pushData(i int){
	s.items =  append(s.items,i)
}

func (s Stack2) GetAll() []int{
	return s.items
}

func main(){
	s:= Stack2{}
	s.pushData(1)
	s.pushData(2)
	s.pushData(3)
	s.pushData(4)

	for _,v1 := range s.GetAll(){
		fmt.Println(v1)
	}
}

