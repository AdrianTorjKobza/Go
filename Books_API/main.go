package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json: "id"`
	Title    string `json: "title"`
	Author   string `json: "author"`
	Year     int    `json: "year"`
	Quantity int    `json: "quantity"`
}

var books = []book{
	{ID: "1", Title: "Book Title 1", Author: "Author 1", Year: 1990, Quantity: 2},
	{ID: "2", Title: "Book Title 2", Author: "Author 2", Year: 2010, Quantity: 10},
	{ID: "3", Title: "Book Title 3", Author: "Author 3", Year: 1981, Quantity: 5},
	{ID: "4", Title: "Book Title 4", Author: "Author 4", Year: 1965, Quantity: 23},
}

// Return the books in json format.
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func bookByID(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Book not found.",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func getBookByID(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("...the book was not found...")
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Missing ID query parameter.",
		})
		return
	}

	book, err := getBookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Book not found.",
		})
		return
	}

	if book.Quantity == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Book not available.",
		})
		return
	}

	book.Quantity = book.Quantity - 1

	c.IndentedJSON(http.StatusOK, book)
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// Add the new book.
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookByID)
	router.POST("/books", createBook)
	router.PATCH("/checkout", checkoutBook)
	router.Run("localhost:8080")
}
