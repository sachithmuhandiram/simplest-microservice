package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting main module")

	http.HandleFunc("/add", add)
	http.ListenAndServe(":7171", nil)

}

func add(res http.ResponseWriter, req *http.Request) {

	add, err := http.Get("http://addmodule:7070/add")

	if err != nil {
		log.Println("Couldnt send request to add module", err)
	}

	log.Println(add)
}
