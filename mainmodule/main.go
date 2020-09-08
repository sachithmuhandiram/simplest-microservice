package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type userInput struct {
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
}

func main() {
	log.Println("Starting main module")

	http.HandleFunc("/add", add)
	http.HandleFunc("/sub", sub)
	http.HandleFunc("/div", div)
	http.HandleFunc("/mul", mul)
	http.ListenAndServe(":7171", nil)

}

func add(res http.ResponseWriter, req *http.Request) {
	/*
		This is the place need to fix. GET request cant get strings converted from int
	*/
	body, readErr := ioutil.ReadAll(req.Body)
	if readErr != nil {
		log.Println(readErr)
	}

	userIn := userInput{}
	jsonErr := json.Unmarshal(body, &userIn)
	if jsonErr != nil {
		log.Println(jsonErr)
	}

	// fmt.Println(userIn.Num1)
	// fmt.Println(userIn.Num2)
	num1 := userIn.Num1 //req.FormValue("num1")
	num2 := userIn.Num2 //req.FormValue("num2")

	log.Println("Num1", num1)
	log.Println("Num2", num2)
	numNew1 := strconv.Itoa(num1)
	numNew2 := strconv.Itoa(num2)
	add, err := http.PostForm("http://addmodule:7070/add", url.Values{"num1": {numNew1}, "num2": {numNew2}})

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
