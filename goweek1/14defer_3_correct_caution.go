package main

import (
	"fmt"
)

func main(){

	fmt.Println("counting")
	
	for i := 0 ; i<10; i++{
		defer func(x int){
		fmt.Println(x)
		}(i) //defer is last in first out
	}	
	
	fmt.Println("done")
}