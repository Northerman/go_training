package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	replacer := strings.NewReplacer(",", "", ".", "", ";", "")
	s = replacer.Replace(s)
	words := strings.Fields(s)
	var m map[string]int //key is string value is string
	m = make(map[string]int)
	for _, word := range words {
		_, ok := m[word]
		if ok {
			m[word] = m[word] + 1
		} else {
			m[word] = 1
		}

	}
	return m
}
func main() {
	s := "If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck"
	w := WordCount(s)
	fmt.Println(w)
}
