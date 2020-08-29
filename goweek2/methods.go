package main

import "fmt"

type rectangle struct {
	width  int
	length int
}

func (rec rectangle) area() int { // area is not a method of rectangle
	return rec.width * rec.length
}

func main() {
	rec := rectangle{3, 4}
	a := rec.area()
	fmt.Println(a)

}
