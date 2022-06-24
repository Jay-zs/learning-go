package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Book struct {
	Id            int     `json:"id"`
	Title         string  `json:"title"`
	Author        *Author `json:"-"`
	Publication   string  `json:"publication"`
	PublishedDate string  `json:"published_date"`
}

type Author struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Dob       string `json:"dob"`
	PenName   string `json:"pen_name"`
}

func getBook(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode([]Book{{1, "Jay", nil, "Jay", "11/03/2002"}, {2, "Goyal", nil, "Goyal", "11/03/2002"}})
}

func getBookById(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode(Book{1, "Jay", nil, "", "11/03/2002"})
}

func postBook(response http.ResponseWriter, request *http.Request) {

}

func postAuthor(response http.ResponseWriter, request *http.Request) {

}

func putAuthor(response http.ResponseWriter, request *http.Request) {

}

func putBook(response http.ResponseWriter, request *http.Request) {

}

func deleteAuthor(response http.ResponseWriter, request *http.Request) {

}

func deleteBook(response http.ResponseWriter, request *http.Request) {

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/book", getBook).Methods(http.MethodGet)

	r.HandleFunc("/book/{id}", getBookById).Methods(http.MethodGet)

	r.HandleFunc("/book", postBook).Methods(http.MethodPost)

	r.HandleFunc("/author", postAuthor).Methods(http.MethodPost)

	r.HandleFunc("/book/{id}", putBook).Methods(http.MethodPut)

	r.HandleFunc("/author/{id}", putAuthor).Methods(http.MethodPut)

	r.HandleFunc("/book/{id}", deleteBook).Methods(http.MethodDelete)

	r.HandleFunc("/author/{id}", deleteAuthor).Methods(http.MethodDelete)

	Server := http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	fmt.Println("Server started at 8000")
	Server.ListenAndServe()
}
