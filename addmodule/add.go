package main

import (
	"log"
	"net/http"
)

func add(reswt http.ResponseWriter, req *http.Request) {
	log.Println("Request came to here")
}

func main() {
	http.HandleFunc("/add", add)
	http.ListenAndServe("0.0.0.0:7070", nil)

}
