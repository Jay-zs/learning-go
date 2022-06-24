package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./")))
	server := http.Server{Addr: ":8000", Handler: mux}
	fmt.Println("Server Listening on 8000")
	server.ListenAndServe()
}
