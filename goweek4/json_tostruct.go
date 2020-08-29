package main

import (
	"encoding/json"
	"fmt"
)

type todo struct {
	ID     int    `json:"id"`
	title  string `json"title"` //note that titleis private cant access
	Status string `json"status"`
}

func main() {
	data := []byte(`{
		"id": 1,
		"title": "pay credit card",
		"status": "active" 
	}`)
	var t todo
	err := json.Unmarshal(data, &t) //we send the pointer so package json knows where is our t
	fmt.Printf("%#v\n", t)
	fmt.Println(err)

}
