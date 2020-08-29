package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func swap(x string, y string) (string,string){
	return y,x
}

func main(){
	fmt.Println(add(12,23))
	fmt.Println(swap("World","Hello"))
	
	// first,second = swap("World","Hello") <-- This cant because have to define variable first
	first,second := swap("World","Hello")
	fmt.Println(first,second)


}