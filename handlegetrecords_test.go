package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MyHandler struct {
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handleGetRecords(w, r)
}

func TestMyHandler(t *testing.T) {
	handler := &MyHandler{}
	server := httptest.NewServer(handler)
	defer server.Close()

	req := httptest.NewRequest("GET", "http://localhost:8084/records/predictions/prediction/5", nil)
	w := httptest.NewRecorder()
	handleGetRecords(w, req)

	resp := w.Result()

	if resp.StatusCode != 200 {
		t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
	}
	expected := "change_me"
	actual, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if expected != string(actual) {
		t.Errorf("Expected the message '%s'\n", expected)
	}
}
