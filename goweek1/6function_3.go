package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(5, 3) * 2

}

func main() {
	// it's called higher order function (function ที่รับหรือ return function)
	hypot := func(x, y float64) float64 {
		return math.Pow(x, y)
	}
	r := hypot(10, 4)
	fmt.Println(r)
	fmt.Println(compute(hypot))
}
