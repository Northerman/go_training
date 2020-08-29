package main

import "fmt"

type Day int

func (d Day) Today() string { // method today ที่คืนค่า string
	return fmt.Sprintf("today: %d", d)

}

type Samsung struct {
	Version string
}

func (s Samsung) Info() string {
	return fmt.Sprintf("info: %s", s.Version)
}

type OnePlus struct {
	Samsung //this is called embed --> oneplus can now use .Info
	// s Samsung // this is not embed --> this is declaring variable name s
	Version string // If we create the same name --> it will be under oneplus
}

func main() {
	d := Day(2)
	today := d.Today()
	fmt.Println(today)

	s := Samsung{}
	s.Version = "s10+"
	fmt.Println(s.Info())
	fmt.Printf("s: %#v\n", s)

	// Note that composite is not inheritance!
	o := OnePlus{}
	o.Samsung.Version = "ของ ss : 7 Pro"
	o.Version = "ของวัน one plus"

	fmt.Println(o.Version)
	fmt.Println(o.Info())
	fmt.Printf("%#v\n", o)
}
