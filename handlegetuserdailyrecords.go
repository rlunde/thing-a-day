package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//router.HandleFunc("/records/{user}/{category}/{field}/{day}/{number}", handleGetUserDailyRecords)
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
	records, err := GetUserDailyRecords(collection, fieldName, user, day)
	if records == nil || err != nil {
		//This gets messy. We don't want to return "recently used" records.
		records, err = GetRandRecordsForUser(collection, fieldName, user, day, nr)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		SaveUserDailyRecords(collection, fieldName, user, day, records)
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
