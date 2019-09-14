package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting main module")

	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/add", add)
	err := http.ListenAndServe(":7171", nil)

	log.Println("Main module web server response : ", err)

}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	log.Println("Test success")
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func add(res http.ResponseWriter, req *http.Request) {

	add, err := http.Get("http://addmodule:7070/add")

	if err != nil {
		log.Println("Couldnt send request to add module", err)
	}

	log.Println(add)
}
