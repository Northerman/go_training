package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	resp := []byte(`{"name": "HA Northerman"}`)
	w.Header().Add("content-type", "application/json")
	w.Header().Add("northerman", "555")
	w.Write(resp)
}
func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("starting")
	// err := http.ListenAndServe(":1234567898", nil) // for invalid port
	err := http.ListenAndServe(":1234", nil) // this is ok!
	log.Fatal(err)
	fmt.Println("end")
	// try open localhost:1234
}
