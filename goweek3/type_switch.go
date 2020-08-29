package main

import "fmt"

type Samsung struct {
	Version string
}

func main() {
	//var i interface{} = "hello"
	var i interface{} = Samsung{Version: "s10+"}
	//k := i.(type) Cant do this (only swtich case can do)
	//fmt.Println(k)
	switch v := i.(type) {
	case int:
		fmt.Printf("%T %d", v, v)
	case string:
		fmt.Printf("%T %s", v, v)
	case Samsung, *Samsung:
		fmt.Printf("%T %#v Yay my new phone \n", v, v)
	default:
		fmt.Println("undefined type")
	}
}
