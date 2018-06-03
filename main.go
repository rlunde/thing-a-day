package main

import (
	"fmt"
	"log"
	"net/http"

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
func makeRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ping", ping)
	router.HandleFunc("/", indexPage).Methods("GET")
	router.HandleFunc("/records/{category}/{field}/{number}", handleGetRecords)
	router.HandleFunc("/records/{user}/{category}/{field}/{day}/{number}", handleGetUserDailyRecords)
	return router
}

/*RunService runs the main service endpoints
 */
func RunService() {
	// TODO: everything
	router := makeRouter()
	log.Fatal(http.ListenAndServe(":8084", router))

}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Thing-A-Day")
	fmt.Println("Endpoint Hit: indexPage")
}
