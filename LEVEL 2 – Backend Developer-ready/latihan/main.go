
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()

    // Routes
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "pong"})
    })

    r.GET("/books", getBooks)
    r.POST("/books", createBook)

    r.Run() // listen and serve on 0.0.0.0:8080
}

type Book struct {
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

var books = []Book{
    {ID: 1, Title: "Go in Action", Author: "William Kennedy"},
    {ID: 2, Title: "The Go Programming Language", Author: "Alan Donovan"},
}

func getBooks(c *gin.Context) {
    c.JSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
    var newBook Book
    if err := c.BindJSON(&newBook); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    newBook.ID = len(books) + 1
    books = append(books, newBook)
    c.JSON(http.StatusCreated, newBook)
}
