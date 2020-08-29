package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./src/basicgo/array.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Println("good bye")

}
