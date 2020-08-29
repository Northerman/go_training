package main

import (
	"fmt"
	"math"
)

const Pi = 3.14 //public variable capital first letter

const PI = 3.14          // Dont do this
const PORT_NUMBER = 8008 // Dont do this
const portNumber = 9009  // GOOD! (camelCase) private variable

const (
	Monday = iota //iota is a generator of running number 1,2,3,..etc
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	_ = iota * 2 //iota is a generator of running number 0,1,2,3,..etc
	First
	Second
)

const (
	Zero = 1024 << iota // "a<<b is a x 2^b    a>>b is a x 2^(-b)"
	One
	Two
)

func main() {
	var x, y int = 3, 4
	// See math.Square https://godoc.org/math
	var f float64 = math.Sqrt(float64(x*x + y*y))
	// var f int32 = math.Sqrt(int32(x*x + y*y))  <-- cant! use int for math.Sqrt
	var z uint = uint(f)
	fmt.Println(x, y, z)

	v := 42.13 //change me!
	k := 42    //change me!
	fmt.Printf("v is of type %T\n", v)
	fmt.Printf("k is of type %T\n", k)

	fmt.Println("Happy", Pi, "Day")
	const truth = true
	fmt.Println("Go rules?", truth)

	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
	fmt.Println(First, Second)
	fmt.Println(Zero, One, Two)
}
