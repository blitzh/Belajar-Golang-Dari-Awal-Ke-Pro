
package main

import "fmt"

func modifySlice(s []int) {
    s[0] = 42
}

func main() {
    original := []int{1, 2, 3}
    modifySlice(original)
    fmt.Println("Modified slice:", original) // [42 2 3]
}
