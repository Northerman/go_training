package main

import (
	"fmt"
	"runtime"
)

func main(){
	os := runtime.GOOS
	// it automatically breaks if it falls in a case
	// if you want to run through use fallthrough
	switch os {
	case "linux":
		fmt.Println("it's linux")
	case "darwin":
		fmt.Println("it's darwin")
	default:
		fmt.Printf("my os %s\n",os)
	
	}

	
}