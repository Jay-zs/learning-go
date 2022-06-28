package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Book struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Author        Author `json:"author"`
	Publication   string `json:"publication"`
	PublishedDate string `json:"published_date"`
}

type Author struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Dob       string `json:"dob"`
	PenName   string `json:"pen_name"`
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "raramuri"
	dbPass := "Goyal@921"
	dbName := "test"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:3306)/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func getBook(response http.ResponseWriter, request *http.Request) {
	db := dbConn()
	defer db.Close()
	title := request.URL.Query().Get("title")
	includeAuthor := request.URL.Query().Get("includeAuthor")
	var rows *sql.Rows
	var err error
	if title == "" {
		rows, err = db.Query("select * from Books;")
	} else {
		rows, err = db.Query("select * from Books where title=?;", title)
	}
	if err != nil {
		log.Print(err)
	}
	books := []Book{}
	for rows.Next() {
		book := Book{}
		err = rows.Scan(&book.Id, &book.Title, &book.Publication, &book.PublishedDate, &book.Author.Id)
		if err != nil {
			log.Print(err)
		}
		if includeAuthor == "true" {
			row := db.QueryRow("select * from Authors where id=?", book.Author.Id)
			row.Scan(&book.Author.Id, &book.Author.FirstName, &book.Author.LastName, &book.Author.Dob, &book.Author.PenName)
		}
		books = append(books, book)
	}
	json.NewEncoder(response).Encode(books)
}

func getBookById(response http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Print(err)
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(Book{})
		return
	}
	db := dbConn()
	defer db.Close()
	bookrow := db.QueryRow("select * from Books where id=?;", id)
	book := Book{}
	err = bookrow.Scan(&book.Id, &book.Title, &book.Publication, &book.PublishedDate, &book.Author.Id)
	if err != nil {
		log.Print(err)
		if err == sql.ErrNoRows {
			response.WriteHeader(404)
			json.NewEncoder(response).Encode(book)
			return
		}
	}
	authorrow := db.QueryRow("select * from Authors where id=?;", book.Author.Id)
	err = authorrow.Scan(&book.Author.Id, &book.Author.FirstName, &book.Author.LastName, &book.Author.Dob, &book.Author.PenName)
	if err != nil {
		log.Print(err)
	}
	json.NewEncoder(response).Encode(book)
}

func postBook(response http.ResponseWriter, request *http.Request) {
	db := dbConn()
	defer db.Close()
	decoder := json.NewDecoder(request.Body)
	b := Book{}
	err := decoder.Decode(&b)
	if b.Title == "" {
		response.WriteHeader(400)
		json.NewEncoder(response).Encode(Book{})
		return
	}
	BookId := 0
	err = db.QueryRow("select id from Books where title=? and author_id=?;", b.Title, b.Author.Id).Scan(&BookId)
	if err == nil {
		response.WriteHeader(400)
		json.NewEncoder(response).Encode(Book{})
		return
	}
	authorRow := db.QueryRow("select id from Authors where id=?;", b.Author.Id)
	authorId := 0
	err = authorRow.Scan(&authorId)
	if err != nil {
		log.Print(err)
		response.WriteHeader(400)
		json.NewEncoder(response).Encode(Book{})
		return
	}
	if !(b.Publication == "Scholastic" || b.Publication == "Penguin" || b.Publication == "Arihanth") {
		response.WriteHeader(400)
		json.NewEncoder(response).Encode(Book{})
		return
	}
	publicationYear, err := strconv.Atoi(strings.Split(b.PublishedDate, "/")[2])
	if err != nil {
		log.Print("invalid date")
		json.NewEncoder(response).Encode(Book{})
		return
	}
	if !(publicationYear >= 1880 && publicationYear <= time.Now().Year()) {
		log.Print("invalid date")
		json.NewEncoder(response).Encode(Book{})
		return
	}
	res, err := db.Exec("INSERT INTO Books (title, publication, published_date, author_id)\nVALUES (?,?,?,?);", b.Title, b.Publication, b.PublishedDate, b.Author.Id)
	id, _ := res.LastInsertId()
	if err != nil {
		log.Print(err)
		json.NewEncoder(response).Encode(Book{})
	} else {
		b.Id = int(id)
		json.NewEncoder(response).Encode(b)
	}
}

func postAuthor(response http.ResponseWriter, request *http.Request) {
	db := dbConn()
	defer db.Close()
	decoder := json.NewDecoder(request.Body)
	a := Author{}
	err := decoder.Decode(&a)
	if a.FirstName == "" || a.Dob == "" {
		response.WriteHeader(400)
		json.NewEncoder(response).Encode(Author{})
		return
	}
	existingAuthorId := 0
	err = db.QueryRow("SELECT id from Authors where first_name=? and last_name=? and dob=? and pen_name=?", a.FirstName, a.LastName, a.Dob, a.PenName).Scan(&existingAuthorId)
	if err == nil {
		log.Print("author already exists")
		response.WriteHeader(400)
		json.NewEncoder(response).Encode(Author{})
		return
	}
	res, err := db.Exec("INSERT INTO Authors (first_name, last_name, dob, pen_name)\nVALUES (?,?,?,?);", a.FirstName, a.LastName, a.Dob, a.PenName)
	id, err := res.LastInsertId()
	if err != nil {
		log.Print(err)
		response.WriteHeader(400)
		json.NewEncoder(response).Encode(Author{})
	} else {
		a.Id = int(id)
		json.NewEncoder(response).Encode(a)
	}
}

func putAuthor(response http.ResponseWriter, request *http.Request) {
	db := dbConn()
	var author Author
	body, err := io.ReadAll(request.Body)
	if err != nil {
		log.Print(err)
		return
	}
	err = json.Unmarshal(body, &author)
	if err != nil {
		log.Print(err)
		return
	}
	if author.FirstName == "" || author.LastName == "" || author.PenName == "" || author.Dob == "" {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	params := mux.Vars(request)
	ID, err := strconv.Atoi(params["id"])
	if ID <= 0 {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := db.Query("SELECT id FROM Authors WHERE id = ?", ID)
	if err != nil {
		log.Print(err)
	}
	if !res.Next() {
		log.Print("id not present")
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	var id int
	err = res.Scan(&id)
	if err != nil {
		log.Print(err)
		return
	}

	_, err = db.Exec("UPDATE Authors SET first_name = ? ,last_name = ? ,dob = ? ,pen_name = ?  WHERE id =?", author.FirstName, author.LastName, author.Dob, author.PenName, ID)
	if err != nil {
		log.Print(err)
		return
	}
	response.WriteHeader(http.StatusOK)
}

func putBook(response http.ResponseWriter, request *http.Request) {

}

func deleteAuthor(response http.ResponseWriter, request *http.Request) {
	db := dbConn()
	defer db.Close()
	id := mux.Vars(request)["id"]
	fmt.Println(id)
	authorId := 0
	exist := db.QueryRow("select id from Authors where id=?;", id)
	err := exist.Scan(&authorId)
	if err == sql.ErrNoRows {
		log.Print(err)
		response.WriteHeader(400)
		return
	} else {
		_, err := db.Exec("delete from Books where author_id=?;", id)
		if err != nil {
			log.Print(err)
			response.WriteHeader(400)
			return
		}
	}
	_, err = db.Exec("delete from Authors where id=?;", id)
	if err != nil {
		response.WriteHeader(400)
		return
	}
	response.WriteHeader(200)
}

func deleteBook(response http.ResponseWriter, request *http.Request) {
	db := dbConn()
	defer db.Close()
	id := mux.Vars(request)["id"]
	bookId := 0
	exist := db.QueryRow("select id from Books where id=?;", id)
	err := exist.Scan(&bookId)
	if err == sql.ErrNoRows {
		log.Print(err)
		response.WriteHeader(400)
		return
	} else {
		_, err = db.Exec("delete from Books where id=?;", id)
		if err != nil {
			response.WriteHeader(400)
			return
		}
	}

	response.WriteHeader(200)
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
