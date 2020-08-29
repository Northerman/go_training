package main

import (
	"errors"
	"fmt"
)

type BusinessError struct {
	Err      error
	Code     int
	Severity int
}

// b := BusinessError{}
// Method we get --> North().
// b := &BusinessError{}
// Method we get --> North(), Error()

func (b BusinessError) North() {
	//....
}

func (b *BusinessError) Error() string { //create a Error method
	return fmt.Sprintf("err=%s, code=%d,severity=%d", b.Err.Error(), b.Code, b.Severity)
}

func PrintErrorMsg(err error) {
	fmt.Println(err)
}

func main() {
	var err error = errors.New("error lol: I am an error") //  error starts with lower case letter and seperate by colon
	PrintErrorMsg(err)

	berr := &BusinessError{Err: err, Code: 555, Severity: 8}
	PrintErrorMsg(berr)


	}

}
