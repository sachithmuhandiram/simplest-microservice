package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
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
	//num2 := req.FormValue("num2")

	log.Println("Data type of the request : ", reflect.TypeOf(num1))
	log.Println("Num1 value : ", num1)
	n1, err := strconv.Atoi(num1)
	//n2, err := strconv.Atoi(num2)

	log.Println("Data type after Atoi : ", reflect.TypeOf(n1))
	if err != nil {
		http.Error(res, "Not a numbers", http.StatusBadRequest)
		return
	}

	fmt.Fprint(res, n1+5)
}
