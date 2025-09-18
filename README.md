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
