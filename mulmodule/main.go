package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Started Mulmodule")

	http.HandleFunc("/mul", mul)
	err := http.ListenAndServe("0.0.0.0:7073", nil)

	log.Println("Mulmodule couldnt start ", err)

}

func mul(res http.ResponseWriter, req *http.Request) {
	log.Println("Request came to Mulmodule container") // just for verification
}
