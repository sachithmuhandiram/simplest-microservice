package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("From main package")

	res, err := http.Get("http://addmodule:7070/add")

	if err != nil {
		log.Println("couldnt send get request")
	}

	log.Println(res)

}
