package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func handleGetUserDailyRecords(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	collection := vars["category"] // the collection of records
	fieldName := vars["field"]     // the collection of records
	numRecords := vars["number"]   // the number to return
	user := vars["user"]
	day := vars["day"]
	nr, nrerr := strconv.Atoi(numRecords)
	if nrerr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//first, see if we already have the records
	records, err := GetUserDailyRecords(collection, fieldname, user, day, nr)
	if records == nil || err != nil {
		records, err = GetRandRecords(collection, fieldName, nr)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		SaveUserDailyRecords(collection, fieldname, user, day, records)
	}
	j, jerr := json.Marshal(records)
	if jerr != nil {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}

}
