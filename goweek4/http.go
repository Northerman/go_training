package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos []Todo

func todosHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(w, "error : %v", err)
			return
		}
		t := Todo{}
		err = json.Unmarshal(body, &t)
		if err != nil {
			fmt.Fprintf(w, "errorasdfadsf: ", err)
		}
		todos = append(todos, t)
		fmt.Printf("body : % #v\n", todos)
		fmt.Fprintf(w, "hello %s created todos", body)
		return
	}
	resp, err := json.Marshal(todos)
	if err != nil {
		fmt.Fprintf(w, "error: ", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
func main() {
	http.HandleFunc("/todos", todosHandler)
	fmt.Println("starting...")
	log.Fatal(http.ListenAndServe(":1234", nil))
	fmt.Println("end")
}
