package main

import (
	"log"
	"net/http"
)

func add(res http.ResponseWriter, req *http.Request) {
	log.Println("Request came to addmodule container")
}

func main() {
	log.Println("Started addmodule")
	http.HandleFunc("/add", add)
	err := http.ListenAndServe("0.0.0.0:7070", nil)

	log.Println("addmodule start values ", err)

}
