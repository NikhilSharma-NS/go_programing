package main

import "fmt"

func Print[L any, V fmt.Stringer](labels []L, vals []V) {
	for i, v := range vals {
		fmt.Println(labels[i], v.String())
	}
}
