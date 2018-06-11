package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type GUDRHandler struct {
}

func (h *GUDRHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handleGetRecords(w, r)
}

//router.HandleFunc("/records/{user}/{category}/{field}/{day}/{number}", handleGetUserDailyRecords)
func TestGUDRHandler(t *testing.T) {
	handler := &GUDRHandler{}
	server := httptest.NewServer(handler)
	defer server.Close()

	req := httptest.NewRequest("GET", "http://localhost:8084/records/humorscope/predictions/prediction/2018-05-22/12", nil)
	w := httptest.NewRecorder()
	handleGetUserDailyRecords(w, req)

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

//router.HandleFunc("/records/{user}/{category}/{field}/{day}/{number}", handleGetUserDailyRecords)
func TestGUDRHandlerWithNames(t *testing.T) {
	handler := &GUDRHandler{}
	server := httptest.NewServer(handler)
	defer server.Close()
	r1 := "http://localhost:8084/records/humorscope/predictions/prediction/2018-05-22/12"
	req := httptest.NewRequest("GET", "http://localhost:8084/records/humorscope/predictions/prediction/2018-05-22/12", nil)
	w := httptest.NewRecorder()
	handleGetUserDailyRecords(w, req)

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
