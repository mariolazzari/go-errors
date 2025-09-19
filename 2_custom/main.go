package main

import "fmt"

type EvenNumberError struct {
	Number int
}

func (e *EvenNumberError) Error() string {
	return fmt.Sprintf("%d is not an odd number", e.Number)
}

func isOdd(n int) error {
	if n%2 == 0 {
		return &EvenNumberError{Number: n}
	}
	return nil
}

func main() {
	var result string
	for i := range 5 {
		err := isOdd(i)
		if err != nil {
			result += fmt.Sprintf("Error: %v\n", err)
		} else {
			result += "Number is odd\n"
		}
	}
}
