package main

import (
	"errors"
	"fmt"
)

func riskyOperation() error {
	return errors.New("something went wrong")
}

func main() {
	err := riskyOperation()
	if err != nil {
		// Proper error handling
		fmt.Printf("Operation failed: %v\n", err)
		return
	}
	fmt.Println("Operation succeeded")
}
