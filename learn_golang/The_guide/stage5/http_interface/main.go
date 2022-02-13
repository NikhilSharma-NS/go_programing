package main

import(
	"fmt"
	"net/http"
	
) 

func main(){

	
	//resp,err:=http.Get("http://google.com")
	fmt.Println("hi")
	callhttp()

}

func callhttp()error{
	resp,err:=http.Get("http://google.com1")
	if err!=nil{
		fmt.Println(err)
		return err
	}
	fmt.Println(resp.Status)
	return nil
}
