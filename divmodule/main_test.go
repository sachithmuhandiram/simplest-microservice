package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestDiv(t *testing.T) {

	req, resErr := http.NewRequest("GET", "http://divmodule:7072/div?num1=11&num2=2", nil)

	if resErr != nil {
		t.Log("NewRequest couldnt handle GET to div module ")
	}

	w := httptest.NewRecorder()

	div(w, req)

	res := w.Result() // response from Add function

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected result ok. got %v from div module ", res.Status)
	}

	bs, err := ioutil.ReadAll(res.Body)

	defer res.Body.Close()

	if err != nil {
		t.Fatalf("Could not read response from div module: %v", err)
	}

	data, err := strconv.Atoi(string(bs))

	if err != nil {
		t.Fatalf("Failed to convert BS to int : %v", err)
	}

	if data != 5 {
		t.Fatalf("Expected 5, got %v", data)
	}

}
