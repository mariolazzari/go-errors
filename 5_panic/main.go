package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	fmt.Println(isOdd(3))
	fmt.Println("Program continues after panic recovery.")
}

func isOdd(n int) bool {
	panic(fmt.Sprintf("a panic message: %d", n))
}
