package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Book defines
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
	Rating int    `json:"rating"`
}

type Books []Book // Defining our type Book

func main() {
	handleRequests()
}

func homepage(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello There, I'm a Golang app")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true) // creating our mux handler
	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/books", allBooks).Methods("GET")
	myRouter.HandleFunc("/books", postBooks).Methods("POST")
	log.Fatal(http.ListenAndServe(":1029", myRouter)) // Our port request
}

func postBooks(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Testing POST endpoint")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

// Creating our allBooks Endpoint
func allBooks(w http.ResponseWriter, r *http.Request) {
	books := Books{
		Book{Title: "Harry Potter and The chamber of secrets",
			Author: "JK Rowling", Year: 1998, Rating: 9},
	}

	fmt.Println("All Books Endpoint")
	err := json.NewEncoder(w).Encode(books) // Here we encode our array 'books'
	if err != nil {
		return
	}
}
