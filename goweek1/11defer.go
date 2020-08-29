package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Println("hello")
	defer fmt.Println("world") // execute after the function ends

	f, err := os.Open("filename.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() //อย่าเพิ่งปิดตอนนี้ ปิดหลังจากทำทุกอย่างเสด (หาง่ายขึ้นมาก)
}
