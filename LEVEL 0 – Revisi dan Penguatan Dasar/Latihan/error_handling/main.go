
package main

import (
    "fmt"
    "os"
)

func openFile(filename string) error {
    f, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("failed to open file: %w", err)
    }
    defer f.Close()
    return nil
}

func main() {
    if err := openFile("missing.txt"); err != nil {
        fmt.Println("Error:", err)
    }
}
