package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("data=%v len=%d cap=%d\n", s, len(s), cap(s))

}

func main() {
	s := []int{}
	printSlice(s)
	if s == nil {
		fmt.Println("s is nil")
	} else {
		fmt.Println("s is", s)
	}

	s = make([]int, 5, 10) //len 5 cap 10

	x := make([]int, 5, 10)
	printSlice(x)
	if x == nil {
		fmt.Println("x is nil")
	} else {
		fmt.Println("x is", x)
	}

}
