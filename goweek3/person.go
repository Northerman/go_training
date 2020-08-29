package main

import "fmt"

type Person struct {
	name    string
	friends map[string]int
}

func (per Person) Walk() string {
	return "Walk " + per.name
}

func (per Person) Eat() string {
	return "Eat " + per.name
}

func (per Person) Greeting() string {
	return "Hello " + per.name
}

func (per Person) Name() string {
	return per.name
}

func (per Person) SetName(n string) {
	fmt.Println("Before:", per.name)
	per.name = n
	per.friends["bez"] = 5                     //eventhough this doesnt use pointer but this changes (it's a sematic pointer(can change w/o pointer) vs sematic value(can't change w/o pointer)..)
	per.friends = map[string]int{"North": 100} //this doesnt change lol... because its a sematic value
	fmt.Println("After:", per.name)

}

func (per *Person) SetNameChange(n string) {
	fmt.Println("Before Inside:", per.name)
	per.name = n
	per.friends["bez"] = 5
	fmt.Println("After Inside:", per.name)

}

func main() {
	per := Person{name: "North",
		friends: map[string]int{
			"bez": 1,
		}}
	fmt.Println(per.Walk())
	fmt.Println(per.Eat())
	fmt.Println(per.Greeting())
	fmt.Println(per.Name()) //Note that if person is public we dont need to write GetName
	fmt.Println(per.name)

	fmt.Println("Before:", per.Name())
	per.SetName("KBTG")
	fmt.Println("After:", per.Name()) // this doesny change because SetName copies "per" not use "per" directly

	fmt.Println("Before Outside:", per.Name())
	per.SetNameChange("KBTG") //use pointer
	fmt.Println("After Outside:", per.Name())

	fmt.Println("Before Outside:", per)
	per.SetName("KBTGv2") //observe the map it can be seen that bez:100 doesnt come out but instead bez 5
	fmt.Println("After Outside:", per)

}
