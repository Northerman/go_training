package main

import (
	"fmt"
	"math"
)

var c, python, java bool = false, false, false
var b = false       // no need to tell type
var x, y = true, 15 // OK!
// var k,v bool = true,14 <-- This is not ok!
// dog := 'I am a dog' <-- cannot use outside of function

func main() {
	var i int = 12
	var java int //it shadows out the var at the top (now java is an int)
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(9))
	fmt.Println("Now you have", math.Sqrt(9), "problems.")
	fmt.Println(c)
	fmt.Println(java)
	fmt.Println(c, java, python)
	fmt.Println(i)
	fmt.Println(b)
	dog := "I am a dog" // <--- declare without using var ( := )
	fmt.Println(dog)
}
