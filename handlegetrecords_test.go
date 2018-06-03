package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestMyHandler(t *testing.T) {
	StartSession()
	req := httptest.NewRequest("GET", "http://localhost:8084/records/predictions/prediction/5", nil)
	w := httptest.NewRecorder()
	r := makeRouter()
	r.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != 200 {
		t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
	}
	actual, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	result := string(actual)
	if len(result) < 20 {
		t.Errorf("result is shorter than expected")
	}
}
