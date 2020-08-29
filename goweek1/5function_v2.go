package main

import (
	"fmt"
)

// Named return values

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func split_confused(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	if x > 2 {
		x = 4
		return
	}
	y = 4
	return
}

func main() {
	fmt.Println("Hello, playground")
	x, y := split(40)
	fmt.Println(x, y)
	k, v := split_confused(40)
	fmt.Println(k, v)
}
