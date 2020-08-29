package main

import "fmt"

type Student struct {
	Name string
	ID   int
}

func (s Student) String() string { // if go detect method that is named String and returns string println will show this !
	return fmt.Sprintf("Hello %d : %s\n", s.ID, s.Name)
}

func main() {
	s := Student{ID: 35, Name: "North"}
	fmt.Println(s)

}
