package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestAdd(t *testing.T) {
	req, resErr := http.NewRequest("GET", "/add?num1=5", nil)

	if resErr != nil {
		t.Fatalf("Couldnt create GET request : %v", resErr)
	}

	w := httptest.NewRecorder()

	Add(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected result ok. got %v: ", res.Status)
	}

	bs, err := ioutil.ReadAll(res.Body)

	defer res.Body.Close()

	if err != nil {
		t.Fatalf("Could not read response : %v", err)
	}

	data, _ := strconv.Atoi(string(bs))

	// if err != nil {
	// 	t.Fatalf("expected an integer, got : %s", err)
	// }

	if data != 8 {
		t.Fatalf("Expected 9, got %v", data)
	}
}
