package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	s := 1
	for s < 1000 {
		s += s
	}
	fmt.Println(s)

	names := []string{"north", "game", "dog"}
	for idx, name := range names {
		fmt.Println(name)
		fmt.Println(idx)
	}
}
