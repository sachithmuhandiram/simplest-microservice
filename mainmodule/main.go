package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting main module")

	http.HandleFunc("/add", add)
	http.HandleFunc("/sub", sub)
	http.HandleFunc("/div", div)
	http.HandleFunc("/mul", mul)
	http.ListenAndServe(":7171", nil)

}

func add(res http.ResponseWriter, req *http.Request) {

	add, err := http.Get("http://addmodule:7070/add")

	if err != nil {
		log.Println("Couldnt send request to add module", err)
	}

	log.Println(add) // Just to verify we gets a response
}

func sub(res http.ResponseWriter, req *http.Request) {

	sub, err := http.Get("http://submodule:7071/sub")

	if err != nil {
		log.Println("Couldnt send request to Sub module", err)
	}

	log.Println(sub) // Just to verify we gets a response
}

func div(res http.ResponseWriter, req *http.Request) {

	div, err := http.Get("http://divmodule:7072/div")

	if err != nil {
		log.Println("Couldnt send request to Div module", err)
	}

	log.Println(div) // Just to verify we gets a response
}

func mul(res http.ResponseWriter, req *http.Request) {

	mul, err := http.Get("http://mulmodule:7073/mul")

	if err != nil {
		log.Println("Couldnt send request to Mul module", err)
	}

	log.Println(mul) // Just to verify we gets a response
}
