package main

import (
	"fmt"
	"net/http"
)

func requestHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Print("Request: ")
	fmt.Println(request)
	fmt.Fprintf(writer, "Hello world!")
}

func main() {
	fmt.Println("Starting webserver ...")
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8081", nil)
}
