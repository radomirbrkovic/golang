// Handling http requests
package main 

import (
	"encoding/json"
	"log"
	"net/http"
)

func countries(w http.ResponseWriter, r *http.Request)  {
	_countries := getData()
	json.NewEncoder(w).Encode(_countries) 
}

func handleRequests() {
	http.HandleFunc("/", countries)
    log.Fatal(http.ListenAndServe(":8081", nil))
}

