package main

import "fmt"

func main() {
	var i interface{} = "hello"
	s, ok := i.(string) // this is assert not convert
	fmt.Println(ok)
	// x, okk := i.(int)
	// fmt.Println(x,okk)
	fmt.Println(s)
	if s, ok := i.(string); ok {
		fmt.Println(s)
		fmt.Println(ok)
	}

}
