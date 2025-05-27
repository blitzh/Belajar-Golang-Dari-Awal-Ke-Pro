
package main

import (
    "errors"
    "fmt"
)

func doSomething() error {
    return fmt.Errorf("wrapped: %w", errors.New("original error"))
}

func main() {
    err := doSomething()
    if errors.Is(err, errors.New("original error")) {
        fmt.Println("original error detected")
    } else {
        fmt.Println("wrapped error:", err)
    }
}
