package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Desc   string
}

var books = []Book{
	{ID: 1, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Desc: "A novel about the decadence of the Roaring Twenties."},
	{ID: 2, Title: "To Kill a Mockingbird", Author: "Harper Lee", Desc: "A novel about racial injustice and the loss of innocence in the South."},
	{ID: 3, Title: "1984", Author: "George Orwell", Desc: "A dystopian novel about a totalitarian government and the power of language."},
}

var PORT = ":8080"

func main() {

	http.HandleFunc("/books", getBooks)
	http.HandleFunc("/book", createBook)
	http.HandleFunc("/book/", bookHandler)
	fmt.Println("Application is listening on port", PORT)
	http.ListenAndServe(PORT, nil)

}

func getBooks(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	if req.Method == "GET" {
		json.NewEncoder(res).Encode(books)
		return
	}

	http.Error(res, "Invalid Method", http.StatusBadRequest)
}

func createBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	if req.Method == "POST" {
		title := req.FormValue("title")
		author := req.FormValue("author")
		desc := req.FormValue("desc")

		newBook := Book{
			ID:     len(books) + 1,
			Title:  title,
			Author: author,
			Desc:   desc,
		}

		books = append(books, newBook)
		json.NewEncoder(res).Encode(newBook)
		return
	}

	http.Error(res, "Invalid Method", http.StatusBadRequest)
}

func bookHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	urlPathSegments := strings.Split(req.URL.Path, "book/")
	bookID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])

	if err != nil {
		http.Error(res, "Invalid book ID", http.StatusBadRequest)
		return
	}

	if req.Method == "GET" {
		book := getBookByID(bookID)

		if book == nil {
			http.Error(res, "Book not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(res).Encode(book)
		return
	} else if req.Method == "PUT" {
		updateBook(res, req, bookID)
		return
	} else if req.Method == "DELETE" {
		deleteBook(res, bookID)
		return
	}

	http.Error(res, "Invalid Method", http.StatusBadRequest)
}

func getBookByID(id int) *Book {
	for _, book := range books {
		if book.ID == id {
			return &book
		}
	}
	return nil
}

func updateBook(res http.ResponseWriter, req *http.Request, id int) {
	for i, book := range books {
		if book.ID == id {
			title := req.FormValue("title")
			author := req.FormValue("author")
			desc := req.FormValue("desc")

			if title != "" {
				books[i].Title = title
			}

			if author != "" {
				books[i].Author = author
			}

			if desc != "" {
				books[i].Desc = desc
			}

			json.NewEncoder(res).Encode(books[i])
			return
		}
	}

	http.Error(res, "Book not found", http.StatusNotFound)
}

func deleteBook(res http.ResponseWriter, id int) {
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			res.WriteHeader(http.StatusOK)
			return
		}
	}
	http.Error(res, "Book not found", http.StatusNotFound)
}
