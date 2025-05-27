
package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    logFile, err := os.Create("app.log")
    if err != nil {
        log.Fatal("failed to create log file")
    }
    defer logFile.Close()

    log.SetOutput(logFile)
    log.Println("App started")

    configValue := os.Getenv("APP_ENV")
    fmt.Println("Config:", configValue)
}
