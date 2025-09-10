package main

import (
	"errors"
	"fmt"
)

func main() {
	err := dosomething()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("No error")
}

type customError struct {
	code    int
	message string
	err     error
}

// Error returns the error message. Implementing error. Method of error interface.
func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s, %v\n", e.code, e.message, e.err)
}

// function that return a custom error.
//
//	func dosomething() error {
//		return &customError{
//			code:    500,
//			message: "something went wrong",
//		}
//	}
func dosomething() error {
	err := doSomethingElse()
	if err != nil {
		return &customError{
			code:    500,
			message: "something went wrong",
			err:     err,
		}
	}
	return nil

}

func doSomethingElse() error {
	return errors.New("Internal error")
}
