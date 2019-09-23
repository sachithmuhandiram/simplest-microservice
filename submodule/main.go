package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	log.Println("Started Submodule")

	http.HandleFunc("/sub", sub)
	err := http.ListenAndServe("0.0.0.0:7071", nil)

	log.Println("Submodule couldnt start ", err)

}

func sub(res http.ResponseWriter, req *http.Request) {
	log.Println("Request came to submodule container") // just for verification

	num1 := req.FormValue("num1")
	num2 := req.FormValue("num2")

	n1, n1Err := strconv.Atoi(num1)
	n2, n2Err := strconv.Atoi(num2)

	if n1Err != nil || n2Err != nil {
		log.Println("Couldnt get num1 or num2 variables")
		return
	}

	fmt.Fprint(res, n1-n2)
}
