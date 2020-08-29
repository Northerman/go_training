package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"asdf"`
	Status string
	Time   time.Time `json:"time"`
}

func main() {
	t := Todo{
		ID:     1,
		Title:  "pay credit card",
		Status: "completed",
		Time:   time.Now(),
	}
	b, err := json.Marshal(t)
	fmt.Println(b)
	fmt.Printf("type : %T \n", b)
	fmt.Printf("byte : %v \n", b)
	fmt.Printf("string: %s \n", b)
	fmt.Println(err)
}
