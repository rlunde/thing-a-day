package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	RunService() // see authapi.go
}

/*RunService runs the main service endpoints
 */
func RunService() {
	// TODO: everything
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ping", ping)
	router.HandleFunc("/", indexPage).Methods("GET")
	log.Fatal(http.ListenAndServe(":8084", router))

}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Thing-A-Day")
	fmt.Println("Endpoint Hit: indexPage")
}
