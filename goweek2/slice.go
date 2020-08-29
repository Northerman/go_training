package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	s := primes[1:4] //change array type to slice type too!
	fmt.Printf("%#v\n", s)
	fmt.Printf("%#v\n", primes)

	s[0] = 6 // the original primes[1] changes too!
	fmt.Printf("%T\n", primes)
	fmt.Printf("%#v\n", primes)
	fmt.Printf("%T\n", s)
	fmt.Printf("%#v\n", s)

	names := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	// This creates a slice & underlying array and names2 is a slice because every slice needs an underlying array
	// names2 := []string{"North", "Hello"}

	k := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
	}
	fmt.Println(k)

	x := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(x)
	fmt.Println("Capacity", cap(x)) //capacity is how much you can hold /hold more
	fmt.Println("Length", len(x))

	x = x[1:4]
	fmt.Println(x)
	fmt.Println("Capacity", cap(x)) //cap = cap เดิม - i = 6-1 = 5
	fmt.Println("Length", len(x))

	x = x[:2]
	fmt.Println(x)
	fmt.Println("Capacity", cap(x)) //cap = cap-0 = 5
	fmt.Println("Length", len(x))

	x = x[1:]
	fmt.Println(x)
	fmt.Println("Capacity", cap(x)) // cap = cap-1 = 4
	fmt.Println("Length", len(x))

	y := append(x, 99, 34) // to increase number of stuff in slice use append
	fmt.Println(y)
	fmt.Println("Capacity", cap(y))
	fmt.Println("Length", len(y))

}
