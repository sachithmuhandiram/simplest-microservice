package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	log.Println("Started addmodule")

	http.HandleFunc("/add", Add)
	err := http.ListenAndServe("0.0.0.0:7070", nil)

	log.Println("addmodule couldnt start ", err)

}

func Add(res http.ResponseWriter, req *http.Request) {
	log.Println("Request came to addmodule container") // just for verification

	num1 := req.FormValue("num1")
	//num2 := req.FormValue("num2")

	log.Println("Data type of the request :", num1)

	n1, err := strconv.Atoi(num1)
	//n2, err := strconv.Atoi(num2)

	if err != nil {
		http.Error(res, "Not a numbers", http.StatusBadRequest)
		return
	}

	fmt.Fprintln(res, n1+5)
}
