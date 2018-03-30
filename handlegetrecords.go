package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func handleGetRecords(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	collection := vars["category"] // the collection of records
	fieldName := vars["field"]     // the collection of records
	numRecords := vars["number"]   // the number to return
	nr, nrerr := strconv.Atoi(numRecords)
	records, err := GetRandRecords(collection, fieldName, nr)
	if nrerr != nil || err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		j, jerr := json.Marshal(records)
		if jerr != nil {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(j)
		}
	}
}
