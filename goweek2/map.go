package main

import "fmt"

func main() {
	var m map[string]string //key is string value is string
	m = make(map[string]string)
	m["north"] = "northerman"
	fmt.Println(m["north"])

}
