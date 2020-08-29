package main

import "fmt"

func main() {
	i := 42
	fmt.Println("i value", i)
	fmt.Println("Address of i", &i)
	var p *int
	p = &i //pointer that is pointed to &i
	fmt.Println("Pointer p is", p)
	fmt.Println("Value of address p is", *p)

	*p = 55
	fmt.Println("p value:", p)
	fmt.Println("p value inside address:", *p)

	fmt.Println("------------Example 2------------------")

	i, j := 42, 2701
	// & is the address * is the pointer
	p = &i          // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // this is now 21

	p = &j         // point to j
	*p = *p / 37   //divide j through the pointer
	fmt.Println(j) // see the new value of j

}
