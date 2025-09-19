# Go Practice: Error Handling

## Error Handling

### Error wrapping

```go
package main

import "fmt"

func main() {
 var result string
 for i := range 5 {
  err := isOdd(i)
  if err != nil {
   result += fmt.Sprintf("Error: %v\n", err)
  } else {
   result += "Number is odd\n"
  }
  // fmt.Println(result)
 }
}

func isOdd(n int) error {
 if n%2 == 0 {
  return fmt.Errorf("number %d is even", n)
 }
 return nil
}
```

### Custom errors

```go
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
```

### Error logging

```go
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
```

### Errors handling in goroutines

```go
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
```

### Panic and recover

```go
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
```
