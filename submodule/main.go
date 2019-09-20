package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Started Submodule")

	http.HandleFunc("/sub", sub)
	err := http.ListenAndServe("0.0.0.0:7071", nil)

	log.Println("Submodule couldnt start ", err)

}

func sub(res http.ResponseWriter, req *http.Request) {
	log.Println("Request came to submodule container") // just for verification
}
