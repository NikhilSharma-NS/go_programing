package main

// func Add[T ???](v1 T,v2 T) T{
// return v1+v2
// }


type addOnly interface{
	type string , int, int8 , int16
}

func Add[T addOnly](v1 T,v2 T){
	return v1+v2
}

func index[T comparable](list []T,find T)int{
	for i,v:= range T{
		if v==find{
			return i
		}
	}
}