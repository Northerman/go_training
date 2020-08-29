package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	//v := Vertex{}
	//v := Vertex{1, 2}
	//v := Vertex{1}    //Cant
	//V := Vertex{X: 1} //OK X will be 1 and Y will be zero value
	v := Vertex{X: 1, Y: 2}

	fmt.Println(v.X)

	//p := &Vertex{}  OK!
	p := &v

	// v.X = 15
	// (*p).X = 19

	fmt.Printf("%#v\n", v)
	fmt.Printf("%#v\n", v.X)
	fmt.Printf("%#v\n", (*p).X) //hard to write (for go if we know that p is pointer it will auto defererence (see below))
	fmt.Printf("%#v\n", p.X)
}
