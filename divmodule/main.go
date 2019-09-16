package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Started Divmodule")

	http.HandleFunc("/div", div)
	err := http.ListenAndServe("0.0.0.0:7072", nil)

	log.Println("Divmodule couldnt start ", err)

}

func div(res http.ResponseWriter, req *http.Request) {
	log.Println("Request came to Divmodule container") // just for verification
}
