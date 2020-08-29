package main

import "fmt"

type Vertex struct {
	X int
	y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 15 //
	v.y = 3  //
	fmt.Println(v)
	fmt.Printf("%+v\n", v)
	fmt.Printf("%#v", v)
	fmt.Println(v.X)
}
