package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/QP", func(writer http.ResponseWriter, request *http.Request) {
		name := request.FormValue("name")
		age := request.FormValue("age")
		fmt.Fprint(writer, "Hello "+name+" your age is "+age)
	})
	server := http.Server{
		Addr:    ":8000",
		Handler: mux,
	}
	fmt.Println("Server started on port 8000")
	server.ListenAndServe()
}
