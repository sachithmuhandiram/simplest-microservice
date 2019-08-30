package main

import (
	"log"
	"net/http"
)

func add(reswt http.ResponseWriter, req *http.Request) {
	log.Println("Request came to here")
}

func main() {
	log.Println("started addmodule")
	http.HandleFunc("/add", add)
	err := http.ListenAndServe("0.0.0.0:7070", nil)

	log.Println("addmodule start values ", err)

}
