package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	mux := http.NewServeMux()
	func ()
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "Hello "+strings.Split(req.URL.Path, "/")[1])
	})
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: mux,
	}
	fmt.Println("Server started and listening on 8000")
	server.ListenAndServe()
}
