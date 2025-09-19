package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	var result string
	for i := range 5 {
		err := isOdd(i)
		if err != nil {
			result += fmt.Sprintf("%v\n", err)
		} else {
			result += fmt.Sprintf("%d is odd.\n", i)
		}
	}
}

func isOdd(n int) error {
	if n%2 == 0 {
		err := errors.New("Logged error: Even number")
		log.Println(err)
		return err
	}
	return nil
}
