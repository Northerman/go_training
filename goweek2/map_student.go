package main

import "fmt"

type Student struct {
	Name  string
	Class string
}

func (stu Student) info() string {
	return stu.Name + stu.Class
}

func main() {
	n := map[string]Student{}
	n["1"] = Student{
		Name:  "Northerman",
		Class: "M6/13",
	}
	n["2"] = Student{Name: "Hello", Class: "Test Class"}
	s, ok := n["2"]
	fmt.Printf("%#v %#v\n", s, ok)

	delete(n, "2")

	s, ok = n["2"]
	fmt.Printf("%#v %#v\n", s, ok)

	n["3"] = Student{Name: "Hello3", Class: "Test Class3"}

	for k, student := range n {
		fmt.Print(k)
		fmt.Println(student.info()) //method call
	}

}
