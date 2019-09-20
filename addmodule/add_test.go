package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestAdd(t *testing.T) {
	req, resErr := http.NewRequest("GET", "http://addmodule:7070/add?num1=5", nil)

	if resErr != nil {
		t.Fatalf("Couldnt create GET request : %v", resErr)
	}

	w := httptest.NewRecorder()

	add(w, req)

	res := w.Result() // response from Add function

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected result ok. got %v: ", res.Status)
	}

	bs, err := ioutil.ReadAll(res.Body)

	defer res.Body.Close()

	if err != nil {
		t.Fatalf("Could not read response : %v", err)
	}

	data, err := strconv.Atoi(string(bs))

	if err != nil {
		t.Fatalf("Failed to convert BS to int : %v", err)
	}

	if data != 10 {
		t.Fatalf("Expected 10, got %v", data)
	}
}
