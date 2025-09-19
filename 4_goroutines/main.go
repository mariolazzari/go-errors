package mai

import (
	"fmt"
	"time"
)

func main() {
	var result string
	for i := 0; i < 5; i++ {
		result = <-isOdd(i)
		fmt.Println(result)
		time.Sleep(time.Millisecond)
	}
}

// isEven() returns an error if a number is even
func isEven(n int) error {
	if n%2 == 0 {
		return fmt.Errorf("%d is an even number", n)
	}
	return nil
}

// isOdd returns a string channel containing
// whether a value is odd
func isOdd(n int) <-chan string {
	result := make(chan string)
	go func() {
		defer close(result)
		if err := isEven(n); err != nil {
			result <- fmt.Sprintf("Error: %v", err)
			return
		}
		result <- fmt.Sprintf("%d is an odd number", n)
	}()

	return result
}
