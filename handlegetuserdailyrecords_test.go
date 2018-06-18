package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-test/deep"
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

func TestGUDRHandlerWithNamesBasicCache(t *testing.T) {
	StartSession()
	// handler := &GUDRHandler{}
	// server := httptest.NewServer(handler)
	// defer server.Close()
	r1 := "http://localhost:8084/records/test/names/firstname/2018-06-10/8"

	//get it the first time
	req1 := httptest.NewRequest("GET", r1, nil)
	w1 := httptest.NewRecorder()

	r := makeRouter()
	r.ServeHTTP(w1, req1)

	resp1 := w1.Result()

	if resp1.StatusCode != 200 {
		t.Fatalf("Received non-200 response: %d\n", resp1.StatusCode)
	}

	res1, err := ioutil.ReadAll(resp1.Body)
	if err != nil {
		t.Fatal(err)
	}
	//get it a second time and verify that it is identical
	req2 := httptest.NewRequest("GET", r1, nil)
	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, req2)

	resp2 := w2.Result()

	if resp2.StatusCode != 200 {
		t.Fatalf("Received non-200 response: %d\n", resp2.StatusCode)
	}

	res2, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		t.Fatal(err)
	}
	if diff := deep.Equal(res1, res2); diff != nil {
		t.Error(diff)
	}

}

//router.HandleFunc("/records/{user}/{category}/{field}/{day}/{number}", handleGetUserDailyRecords)
func TestGUDRHandlerWithNames(t *testing.T) {
	handler := &GUDRHandler{}
	server := httptest.NewServer(handler)
	defer server.Close()
	r1 := "http://localhost:8084/records/test/names/firstname/2018-06-10/8"
	r2 := "http://localhost:8084/records/test/names/firstname/2018-06-11/8"
	r3 := "http://localhost:8084/records/test/names/firstname/2018-06-12/8"
	r4 := "http://localhost:8084/records/test/names/firstname/2018-06-13/8"

	reqs := [4]string{r1, r2, r3, r4}
	for _, r := range reqs {
		req := httptest.NewRequest("GET", r, nil)
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

}
