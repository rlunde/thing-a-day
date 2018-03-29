package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	err := StartSession()
	defer EndSession()
	if err != nil {
		panic(err)
	}
	RunService() // see authapi.go
}

/*RunService runs the main service endpoints
 */
func RunService() {
	// TODO: everything
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ping", ping)
	router.HandleFunc("/", indexPage).Methods("GET")
	router.HandleFunc("/records/{category}/{field}/{number}", handleGetRecords)
	log.Fatal(http.ListenAndServe(":8084", router))

}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Thing-A-Day")
	fmt.Println("Endpoint Hit: indexPage")
}

func handleGetRecords(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	collection := vars["category"] // the collection of records
	fieldName := vars["field"]     // the collection of records
	numRecords := vars["number"]   // the number to return
	nr, nrerr := strconv.Atoi(numRecords)
	records, err := randRecords(collection, fieldName, nr)
	if nrerr != nil || err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		j, jerr := json.Marshal(records)
		if jerr != nil {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(j)
		}
	}
}
