package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func todosHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(w, "error : %v", err)
			return
		}
		fmt.Printf("body : %s\n", body)
		fmt.Fprintf(w, "hello %s created todos", "POST")
		return
	}
	fmt.Fprintf(w, "hello %s todos", "GET")
}
func main() {
	http.HandleFunc("/todos", todosHandler)
	log.Fatal(http.ListenAndServe(":1234", nil))
}
