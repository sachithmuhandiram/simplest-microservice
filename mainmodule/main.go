package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting main module")
	http.ListenAndServe(":7171", nil)

	http.HandleFunc("/add", add)
	//err := http.ListenAndServe(":7070", nil)

	// if err != nil {
	// 	log.Println("couldnt start web service at main module")
	// }

	//log.Println("Main module web server response : ", err)

}

func add(res http.ResponseWriter, req *http.Request) {

	add, err := http.Get("http://addmodule:7070/add")

	if err != nil {
		log.Println("Couldnt send request to add module")
	}

	log.Println(add)
}
