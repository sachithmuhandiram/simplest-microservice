package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	log.Println("Started addmodule")

	http.HandleFunc("/add", add)
	err := http.ListenAndServe("0.0.0.0:7070", nil)

	log.Println("addmodule couldnt start ", err)

}

func add(res http.ResponseWriter, req *http.Request) {
	log.Println("Request came to addmodule container") // just for verification

	num1 := req.FormValue("num1")
	num2 := req.FormValue("num2")
	log.Println(num1, num2)
	n1, n1Err := strconv.Atoi(num1) // converts string to int
	n2, n2Err := strconv.Atoi(num2)

	if n1Err != nil || n2Err != nil {
		http.Error(res, "Couldnt convert to integers in using Atoi", http.StatusBadRequest)
		return
	}

	log.Println("Answer is : ", n1+n2)
	fmt.Fprint(res, n1+n2)
}
