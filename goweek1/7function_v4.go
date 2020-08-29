package main

import "fmt"

func adder() (func() int, func() int) { //returns two function
	sum := 0
	return func() int { // return func(),func()
			sum += 1
			return sum
		},
		func() int {
			return sum
		}
}

func main() {
	inc, cur := adder()
	fmt.Println(cur())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(cur())

	inc_2, cur_2 := adder()
	fmt.Println(cur_2())
	fmt.Println(inc_2())
	fmt.Println(inc_2())
	fmt.Println(cur_2())

}
