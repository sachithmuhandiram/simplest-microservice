package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestMul(t *testing.T) {

	req, resErr := http.NewRequest("GET", "http://mulmodule:7073/mul?num1=5&num2=2", nil)

	if resErr != nil {
		t.Log("NewRequest couldnt handle GET to mul module ")
	}

	w := httptest.NewRecorder()

	mul(w, req)

	res := w.Result() // response from Mul function

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected result ok. got %v from mul module ", res.Status)
	}

	bs, err := ioutil.ReadAll(res.Body)

	defer res.Body.Close()

	if err != nil {
		t.Fatalf("Could not read response from mul module: %v", err)
	}

	data, err := strconv.Atoi(string(bs))

	if err != nil {
		t.Fatalf("Failed to convert BS to int : %v", err)
	}

	if data != 10 {
		t.Fatalf("Expected 10, got %v", data)
	}
}
