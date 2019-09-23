package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestSub(t *testing.T) {

	req, resErr := http.NewRequest("GET", "http://submodule:7071/sub?num1=9&num2=2", nil)

	if resErr != nil {
		t.Log("NewRequest couldnt handle GET to sub module ")
	}

	w := httptest.NewRecorder()

	sub(w, req)

	res := w.Result() // response from Sub function

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected result ok. got %v from sub module ", res.Status)
	}

	bs, err := ioutil.ReadAll(res.Body)

	defer res.Body.Close()

	if err != nil {
		t.Fatalf("Could not read response from sub module: %v", err)
	}

	data, err := strconv.Atoi(string(bs))

	if err != nil {
		t.Fatalf("Failed to convert BS to int : %v", err)
	}

	if data != 5 {
		t.Fatalf("Expected 5, got %v", data)
	}
}
